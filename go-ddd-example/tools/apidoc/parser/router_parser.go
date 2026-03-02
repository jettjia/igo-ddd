package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type RouterParser struct {
	config *Config
}

func NewRouterParser(config *Config) *RouterParser {
	return &RouterParser{config: config}
}

func (p *RouterParser) Parse() ([]*Route, error) {
	var routes []*Route

	// 解析主路由文件，获取基础前缀
	basePrefix := p.parseMainRouter()

	// 只解析 public 路由，不解析 private
	publicRoutes, err := p.parseRouterDirectory(filepath.Join(p.config.RouterPath, "public"), basePrefix)
	if err != nil {
		return nil, err
	}
	routes = append(routes, publicRoutes...)

	return routes, nil
}

func (p *RouterParser) parseMainRouter() string {
	mainRouterPath := filepath.Join(p.config.RouterPath, "router.go")
	if _, err := os.Stat(mainRouterPath); os.IsNotExist(err) {
		return ""
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, mainRouterPath, nil, parser.ParseComments)
	if err != nil {
		return ""
	}

	var basePrefix string
	ast.Inspect(file, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Group" {
				if len(call.Args) > 0 {
					if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
						basePrefix = strings.Trim(lit.Value, "\"")
					}
				}
			}
		}
		return true
	})

	return basePrefix
}

func (p *RouterParser) parseRouterDirectory(dirPath, basePrefix string) ([]*Route, error) {
	var routes []*Route

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fileRoutes, err := p.parseRouterFile(path, basePrefix)
		if err != nil {
			return err
		}
		routes = append(routes, fileRoutes...)

		return nil
	})

	return routes, err
}

func (p *RouterParser) parseRouterFile(filePath, basePrefix string) ([]*Route, error) {
	var routes []*Route

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 获取包名
	pkgName := file.Name.Name

	// 解析函数调用
	ast.Inspect(file, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok {
			// 查找设置路由的函数
			if strings.HasPrefix(funcDecl.Name.Name, "Set") &&
				(strings.HasSuffix(funcDecl.Name.Name, "Router") ||
					strings.Contains(funcDecl.Name.Name, "Router")) {
				funcRoutes := p.parseRouterFunction(funcDecl, basePrefix, pkgName)
				routes = append(routes, funcRoutes...)
			}
		}
		return true
	})

	return routes, nil
}

func (p *RouterParser) parseRouterFunction(funcDecl *ast.FuncDecl, basePrefix, pkgName string) []*Route {
	var routes []*Route
	var currentGroupPrefix string
	var currentGroupName string

	ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.AssignStmt:
			// 解析 GRouter := Router.Group("/xxx")
			for i, lhs := range node.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok && ident.Name == "GRouter" {
					if i < len(node.Rhs) {
						if call, ok := node.Rhs[i].(*ast.CallExpr); ok {
							if sel, ok := call.Fun.(*ast.SelectorExpr); ok && sel.Sel.Name == "Group" {
								if len(call.Args) > 0 {
									if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
										currentGroupPrefix = strings.Trim(lit.Value, "\"")
										// 提取分组名称，去掉开头的斜杠
										currentGroupName = strings.TrimPrefix(currentGroupPrefix, "/")
									}
								}
							}
						}
					}
				}
			}

		case *ast.ExprStmt:
			// 解析路由注册 GRouter.POST("/xxx", handler)
			if call, ok := node.X.(*ast.CallExpr); ok {
				if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
					method := strings.ToUpper(sel.Sel.Name)
					if isHTTPMethod(method) && len(call.Args) >= 2 {
						// 解析路径
						var path string
						if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
							path = strings.Trim(lit.Value, "\"")
						}

						// 解析 handler
						var handler string
						if sel, ok := call.Args[1].(*ast.SelectorExpr); ok {
							handler = sel.Sel.Name
						}

						if path != "" && handler != "" {
							// 构建完整路径
							fullPath := p.buildFullPath(basePrefix, currentGroupPrefix, path)

							route := &Route{
								Path:       fullPath,
								Method:     method,
								Handler:    handler,
								HandlerPkg: currentGroupName, // 使用分组名称作为 HandlerPkg
							}
							routes = append(routes, route)
						}
					}
				}
			}
		}
		return true
	})

	return routes
}

func (p *RouterParser) buildFullPath(basePrefix, groupPrefix, path string) string {
	var parts []string

	if basePrefix != "" {
		parts = append(parts, strings.TrimPrefix(basePrefix, "/"))
	}

	if groupPrefix != "" {
		parts = append(parts, strings.TrimPrefix(groupPrefix, "/"))
	}

	if path != "" {
		parts = append(parts, strings.TrimPrefix(path, "/"))
	}

	return "/" + strings.Join(parts, "/")
}

func isHTTPMethod(method string) bool {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}
