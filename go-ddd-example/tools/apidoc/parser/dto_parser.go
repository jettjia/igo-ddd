package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type DtoParser struct {
	config *Config
}

func NewDtoParser(config *Config) *DtoParser {
	return &DtoParser{config: config}
}

func (p *DtoParser) Parse() ([]*DtoInfo, error) {
	var dtos []*DtoInfo

	// 遍历 DTO 文件
	err := filepath.Walk(p.config.DtoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// 解析 Go 文件
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		// 获取包名
		pkgName := file.Name.Name

		// 遍历所有类型声明
		ast.Inspect(file, func(n ast.Node) bool {
			if typeSpec, ok := n.(*ast.TypeSpec); ok {
				if structType, ok := typeSpec.Type.(*ast.StructType); ok {
					dto := p.parseStructType(typeSpec.Name.Name, pkgName, structType)
					if dto != nil {
						dtos = append(dtos, dto)
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

	return dtos, nil
}

func (p *DtoParser) parseStructType(name, pkgName string, structType *ast.StructType) *DtoInfo {
	dto := &DtoInfo{
		Name:    name,
		Package: pkgName,
	}

	for _, field := range structType.Fields.List {
		if field.Names == nil {
			continue // 跳过匿名字段
		}

		fieldName := field.Names[0].Name

		// 过滤隐藏字段
		if p.isHiddenField(fieldName) {
			continue
		}

		fieldInfo := &DtoFieldInfo{
			Name: fieldName,
		}

		// 获取字段类型
		fieldInfo.Type = p.getFieldType(field.Type)

		// 获取字段标签
		if field.Tag != nil {
			fieldInfo.Tag = field.Tag.Value
		}

		// 获取字段注释
		if field.Doc != nil {
			fieldInfo.Comment = field.Doc.Text()
		} else if field.Comment != nil {
			fieldInfo.Comment = field.Comment.Text()
		}

		// 检查是否必填
		if field.Tag != nil {
			tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
			if v := tag.Get("validate"); strings.Contains(v, "required") {
				fieldInfo.Required = true
			}
		}

		dto.Fields = append(dto.Fields, fieldInfo)
	}

	return dto
}

// isHiddenField 检查字段是否为隐藏字段
func (p *DtoParser) isHiddenField(fieldName string) bool {
	hiddenFields := []string{
		"CreatedBy",
		"DeletedBy",
		"UpdatedBy",
		"created_by",
		"deleted_by",
		"updated_by",
	}

	for _, hidden := range hiddenFields {
		if fieldName == hidden {
			return true
		}
	}

	return false
}

func (p *DtoParser) getFieldType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name + "." + t.Sel.Name
		}
		return t.Sel.Name
	case *ast.StarExpr:
		return "*" + p.getFieldType(t.X)
	case *ast.ArrayType:
		return "[]" + p.getFieldType(t.Elt)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.MapType:
		return "map[" + p.getFieldType(t.Key) + "]" + p.getFieldType(t.Value)
	default:
		return "unknown"
	}
}
