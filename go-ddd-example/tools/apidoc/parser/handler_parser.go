package parser

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"
)

type HandlerParser struct {
	config *Config
}

func NewHandlerParser(config *Config) *HandlerParser {
	return &HandlerParser{config: config}
}

func (p *HandlerParser) Parse() ([]*Handler, error) {
	var handlers []*Handler

	err := filepath.Walk(p.config.HandlerPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		pkgName := file.Name.Name

		ast.Inspect(file, func(n ast.Node) bool {
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				if funcDecl.Recv != nil {
					handler := p.parseHandlerFunc(funcDecl, pkgName, file, fset)
					if handler != nil {
						handlers = append(handlers, handler)
					}
				}
			}
			return true
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}

func (p *HandlerParser) parseHandlerFunc(funcDecl *ast.FuncDecl, pkgName string, file *ast.File, fset *token.FileSet) *Handler {
	handler := &Handler{
		Name:    funcDecl.Name.Name,
		Package: pkgName,
	}

	// 0. 优先从注释中查找 @request/@response/@summary/@desc
	if funcDecl.Doc != nil {
		for _, comment := range funcDecl.Doc.List {
			text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
			if strings.HasPrefix(text, "@request ") {
				name := strings.TrimSpace(strings.TrimPrefix(text, "@request "))
				if name != "" {
					handler.Request = &DtoInfo{Name: name}
				}
			}
			if strings.HasPrefix(text, "@response ") {
				name := strings.TrimSpace(strings.TrimPrefix(text, "@response "))
				if name != "" {
					handler.Response = &DtoInfo{Name: name}
				}
			}
			if strings.HasPrefix(text, "@summary ") {
				handler.Summary = strings.TrimSpace(strings.TrimPrefix(text, "@summary "))
			}
			if strings.HasPrefix(text, "@desc ") {
				handler.Desc = strings.TrimSpace(strings.TrimPrefix(text, "@desc "))
			}
		}
	}

	// 只有未手动指定时才自动推断
	if handler.Request == nil {
		p.parseFunctionParams(funcDecl, handler)
	}
	if handler.Response == nil {
		p.parseFunctionReturns(funcDecl, handler)
	}
	if handler.Request == nil || handler.Response == nil {
		p.parseFunctionBody(funcDecl, handler, file, fset)
	}

	return handler
}

func (p *HandlerParser) parseFunctionParams(funcDecl *ast.FuncDecl, handler *Handler) {
	if funcDecl.Type.Params == nil {
		return
	}

	for _, param := range funcDecl.Type.Params.List {
		if len(param.Names) == 0 {
			continue
		}

		// 检查是否是 gin.Context 参数
		if sel, ok := param.Type.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == "gin" && sel.Sel.Name == "Context" {
				continue // 跳过 gin.Context 参数
			}
		}

		// 检查是否是请求参数
		for _, name := range param.Names {
			if strings.Contains(strings.ToLower(name.Name), "req") ||
				strings.Contains(strings.ToLower(name.Name), "request") {
				if sel, ok := param.Type.(*ast.SelectorExpr); ok {
					if strings.HasSuffix(sel.Sel.Name, "Req") || strings.HasSuffix(sel.Sel.Name, "Request") {
						handler.Request = &DtoInfo{Name: sel.Sel.Name}
						return
					}
				}
			}
		}
	}
}

func (p *HandlerParser) parseFunctionReturns(funcDecl *ast.FuncDecl, handler *Handler) {
	if funcDecl.Type.Results == nil {
		return
	}

	for _, result := range funcDecl.Type.Results.List {
		if sel, ok := result.Type.(*ast.SelectorExpr); ok {
			if strings.HasSuffix(sel.Sel.Name, "Rsp") ||
				strings.HasSuffix(sel.Sel.Name, "Response") ||
				strings.HasSuffix(sel.Sel.Name, "Res") {
				handler.Response = &DtoInfo{Name: sel.Sel.Name}
				return
			}
		}
	}
}

func (p *HandlerParser) parseFunctionBody(funcDecl *ast.FuncDecl, handler *Handler, file *ast.File, fset *token.FileSet) {
	if funcDecl.Body == nil {
		return
	}

	// 设置类型检查器
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	_, err := conf.Check(handler.Package, fset, []*ast.File{file}, info)
	if err != nil {
		// 类型检查失败时使用 AST 解析
		p.parseFunctionBodyAST(funcDecl, handler)
		return
	}

	// 记录变量名到类型的映射
	varMap := make(map[string]string)

	// 识别变量声明
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.DeclStmt:
			if gen, ok := node.Decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
				for _, spec := range gen.Specs {
					if vs, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range vs.Names {
							if sel, ok := vs.Type.(*ast.SelectorExpr); ok {
								if strings.HasSuffix(sel.Sel.Name, "Req") || strings.HasSuffix(sel.Sel.Name, "Request") {
									varMap[name.Name] = sel.Sel.Name
								}
							}
						}
					}
				}
			}
		case *ast.AssignStmt:
			// 识别 dtoReq := dtoUser.XXXReq{}
			for i, lhs := range node.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && i < len(node.Rhs) {
					if comp, ok := node.Rhs[i].(*ast.CompositeLit); ok {
						if sel, ok := comp.Type.(*ast.SelectorExpr); ok {
							if strings.HasSuffix(sel.Sel.Name, "Req") || strings.HasSuffix(sel.Sel.Name, "Request") {
								varMap[ident.Name] = sel.Sel.Name
							}
						}
					}
				}
			}
		}
		return true
	})

	// 识别 BindJSON 调用
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "BindJSON" {
				if len(call.Args) == 1 {
					if unary, ok := call.Args[0].(*ast.UnaryExpr); ok {
						if ident, ok := unary.X.(*ast.Ident); ok {
							if reqType, ok := varMap[ident.Name]; ok {
								handler.Request = &DtoInfo{Name: reqType}
							}
						}
					}
				}
			}
		}
		return true
	})

	// 自动识别 xresponse.RspOk(c, ..., rsp) 的返回类型
	if handler.Response == nil {
		ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			// 检查是否是 xresponse.RspOk 调用
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "RspOk" {
				if len(call.Args) >= 3 {
					if ident, ok := call.Args[2].(*ast.Ident); ok {
						if obj, ok := info.Uses[ident]; ok {
							if v, ok := obj.(*types.Var); ok {
								typ := v.Type().String()
								parts := strings.Split(typ, ".")
								if len(parts) > 1 {
									lastPart := parts[len(parts)-1]
									handler.Response = &DtoInfo{Name: lastPart}
								}
							}
						}
					}
				}
			}
			return true
		})
	}
}

func (p *HandlerParser) parseFunctionBodyAST(funcDecl *ast.FuncDecl, handler *Handler) {
	// 纯 AST 解析，不依赖类型检查
	varMap := make(map[string]string)

	// 识别变量声明和赋值
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.DeclStmt:
			if gen, ok := node.Decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
				for _, spec := range gen.Specs {
					if vs, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range vs.Names {
							if sel, ok := vs.Type.(*ast.SelectorExpr); ok {
								if strings.HasSuffix(sel.Sel.Name, "Req") || strings.HasSuffix(sel.Sel.Name, "Request") {
									varMap[name.Name] = sel.Sel.Name
								}
							}
						}
					}
				}
			}
		case *ast.AssignStmt:
			for i, lhs := range node.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && i < len(node.Rhs) {
					if comp, ok := node.Rhs[i].(*ast.CompositeLit); ok {
						if sel, ok := comp.Type.(*ast.SelectorExpr); ok {
							if strings.HasSuffix(sel.Sel.Name, "Req") || strings.HasSuffix(sel.Sel.Name, "Request") {
								varMap[ident.Name] = sel.Sel.Name
							}
						}
					}
				}
			}
		}
		return true
	})

	// 识别 BindJSON 调用
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "BindJSON" {
				if len(call.Args) == 1 {
					if unary, ok := call.Args[0].(*ast.UnaryExpr); ok {
						if ident, ok := unary.X.(*ast.Ident); ok {
							if reqType, ok := varMap[ident.Name]; ok {
								handler.Request = &DtoInfo{Name: reqType}
							}
						}
					}
				}
			}
		}
		return true
	})

	// 识别响应变量声明
	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		if assign, ok := n.(*ast.AssignStmt); ok {
			for i, lhs := range assign.Lhs {
				if _, ok := lhs.(*ast.Ident); ok && i < len(assign.Rhs) {
					if comp, ok := assign.Rhs[i].(*ast.CompositeLit); ok {
						if sel, ok := comp.Type.(*ast.SelectorExpr); ok {
							if strings.HasSuffix(sel.Sel.Name, "Rsp") ||
								strings.HasSuffix(sel.Sel.Name, "Response") ||
								strings.HasSuffix(sel.Sel.Name, "Res") {
								handler.Response = &DtoInfo{Name: sel.Sel.Name}
								return false
							}
						}
					}
				}
			}
		}
		return true
	})
}
