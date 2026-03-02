package generator

import (
	"fmt"
	"strings"

	"tools/apidoc/parser"
)

type Config struct {
	Title       string
	Version     string
	Description string
}

type OpenAPIGenerator struct {
	config *Config
}

func NewOpenAPIGenerator(config *Config) *OpenAPIGenerator {
	return &OpenAPIGenerator{config: config}
}

func (g *OpenAPIGenerator) Generate(routes []*parser.Route, handlers []*parser.Handler, dtos []*parser.DtoInfo) map[string]interface{} {
	// OpenAPI 3.0 基本结构
	doc := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       g.config.Title,
			"version":     g.config.Version,
			"description": g.config.Description,
		},
		"paths":      make(map[string]interface{}),
		"components": map[string]interface{}{"schemas": make(map[string]interface{})},
	}

	// 生成 schemas
	schemas := doc["components"].(map[string]interface{})["schemas"].(map[string]interface{})
	// 收集所有 DTO 名称，便于嵌套类型判断
	dtoNames := map[string]struct{}{}
	for _, dto := range dtos {
		dtoNames[dto.Name] = struct{}{}
	}
	for _, dto := range dtos {
		schemas[dto.Name] = g.generateSchemaWithDTOs(dto, dtoNames)
	}

	// 生成 paths
	paths := doc["paths"].(map[string]interface{})
	for _, route := range routes {
		handler := g.findHandler(handlers, route.Handler)
		if handler != nil {
			pathItem := g.generatePathItem(route, handler)
			if existing, ok := paths[route.Path].(map[string]interface{}); ok {
				for k, v := range pathItem {
					existing[k] = v
				}
			} else {
				paths[route.Path] = pathItem
			}
		}
	}

	return doc
}

func (g *OpenAPIGenerator) generateSchema(dto *parser.DtoInfo) map[string]interface{} {
	properties := make(map[string]interface{})
	required := []string{}

	for _, field := range dto.Fields {
		fieldSchema := g.generateFieldSchema(field)

		// 从 json tag 中获取字段名
		jsonName := field.Name
		if field.Tag != "" {
			tag := strings.Trim(field.Tag, "`")
			if jsonTag := strings.Split(tag, "json:\""); len(jsonTag) > 1 {
				jsonName = strings.Split(jsonTag[1], "\"")[0]
			}
		}

		if field.Required {
			required = append(required, jsonName)
		}

		properties[jsonName] = fieldSchema
	}

	schema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

func (g *OpenAPIGenerator) generateFieldSchema(field *parser.DtoFieldInfo) map[string]interface{} {
	schema := make(map[string]interface{})

	// 根据字段类型设置 schema
	switch {
	case strings.HasPrefix(field.Type, "[]"):
		schema["type"] = "array"
		itemType := strings.TrimPrefix(field.Type, "[]")
		if strings.HasPrefix(itemType, "*") {
			itemType = strings.TrimPrefix(itemType, "*")
		}
		// 处理包名前缀，如 builder.Query -> Query
		if strings.Contains(itemType, ".") {
			parts := strings.Split(itemType, ".")
			itemType = parts[len(parts)-1]
		}
		schema["items"] = map[string]interface{}{
			"$ref": fmt.Sprintf("#/components/schemas/%s", itemType),
		}
	case strings.HasPrefix(field.Type, "*"):
		refType := strings.TrimPrefix(field.Type, "*")
		if strings.Contains(refType, ".") {
			parts := strings.Split(refType, ".")
			refType = parts[len(parts)-1]
		}
		schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
	case field.Type == "interface{}":
		schema["type"] = "object"
	case strings.HasPrefix(field.Type, "map["):
		schema["type"] = "object"
		schema["additionalProperties"] = map[string]interface{}{
			"type": "string",
		}
	default:
		// 处理基本类型中的包名前缀
		if strings.Contains(field.Type, ".") {
			parts := strings.Split(field.Type, ".")
			refType := parts[len(parts)-1]
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
		} else {
			schema["type"] = g.mapGoTypeToOpenAPI(field.Type)
		}
	}

	if field.Comment != "" {
		schema["description"] = strings.TrimSpace(field.Comment)
	}

	return schema
}

func (g *OpenAPIGenerator) generatePathItem(route *parser.Route, handler *parser.Handler) map[string]interface{} {
	summary := fmt.Sprintf("%s %s", route.Method, route.Path)
	if handler.Summary != "" {
		summary = handler.Summary
	}
	desc := ""
	if handler.Desc != "" {
		desc = handler.Desc
	}

	operation := map[string]interface{}{
		"summary":     summary,
		"description": desc,
		"operationId": handler.Name,
		"tags":        []string{route.HandlerPkg},
		"responses":   make(map[string]interface{}),
	}

	// 添加请求参数
	if handler.Request != nil {
		operation["requestBody"] = map[string]interface{}{
			"required": true,
			"content": map[string]interface{}{
				"application/json": map[string]interface{}{
					"schema": map[string]interface{}{
						"$ref": fmt.Sprintf("#/components/schemas/%s", handler.Request.Name),
					},
				},
			},
		}
	}

	// 添加 URL 参数
	if len(handler.URLParams) > 0 {
		parameters := []map[string]interface{}{}
		for _, param := range handler.URLParams {
			parameters = append(parameters, map[string]interface{}{
				"name":     param,
				"in":       "path",
				"required": true,
				"schema": map[string]interface{}{
					"type": "string",
				},
			})
		}
		operation["parameters"] = parameters
	}

	// 添加响应
	responses := operation["responses"].(map[string]interface{})
	if handler.Response != nil {
		schema := map[string]interface{}{}
		if strings.HasPrefix(handler.Response.Name, "[]") {
			itemType := strings.TrimPrefix(handler.Response.Name, "[]")
			itemType = strings.TrimPrefix(itemType, "*") // 支持 *[]XXX
			schema["type"] = "array"
			schema["items"] = map[string]interface{}{
				"$ref": fmt.Sprintf("#/components/schemas/%s", itemType),
			}
		} else {
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", handler.Response.Name)
		}
		responses["200"] = map[string]interface{}{
			"description": "Success",
			"content": map[string]interface{}{
				"application/json": map[string]interface{}{
					"schema": schema,
				},
			},
		}
	} else {
		responses["204"] = map[string]interface{}{
			"description": "No Content",
		}
	}

	// 添加错误响应
	responses["400"] = map[string]interface{}{
		"description": "Bad Request",
	}
	responses["500"] = map[string]interface{}{
		"description": "Internal Server Error",
	}

	return map[string]interface{}{
		strings.ToLower(route.Method): operation,
	}
}

func (g *OpenAPIGenerator) findHandler(handlers []*parser.Handler, name string) *parser.Handler {
	for _, h := range handlers {
		if h.Name == name {
			return h
		}
	}
	return nil
}

func (g *OpenAPIGenerator) mapGoTypeToOpenAPI(goType string) string {
	switch goType {
	case "string":
		return "string"
	case "int", "int32", "int64":
		return "integer"
	case "float32", "float64":
		return "number"
	case "bool":
		return "boolean"
	case "time.Time":
		return "string"
	case "builder.Query", "builder.PageData", "builder.SortData":
		return "object"
	default:
		return "string"
	}
}

// generateSchemaWithDTOs 支持嵌套结构体
func (g *OpenAPIGenerator) generateSchemaWithDTOs(dto *parser.DtoInfo, dtoNames map[string]struct{}) map[string]interface{} {
	properties := make(map[string]interface{})
	required := []string{}

	for _, field := range dto.Fields {
		fieldSchema := g.generateFieldSchemaWithDTOs(field, dtoNames)
		// 从 json tag 中获取字段名
		jsonName := field.Name
		if field.Tag != "" {
			tag := strings.Trim(field.Tag, "`")
			if jsonTag := strings.Split(tag, "json:\""); len(jsonTag) > 1 {
				jsonName = strings.Split(jsonTag[1], "\"")[0]
			}
		}

		if field.Required {
			required = append(required, jsonName)
		}

		properties[jsonName] = fieldSchema
	}

	schema := map[string]interface{}{
		"type":       "object",
		"properties": properties,
	}

	if len(required) > 0 {
		schema["required"] = required
	}

	return schema
}

// generateFieldSchemaWithDTOs 支持嵌套结构体 $ref
func (g *OpenAPIGenerator) generateFieldSchemaWithDTOs(field *parser.DtoFieldInfo, dtoNames map[string]struct{}) map[string]interface{} {
	schema := make(map[string]interface{})

	// 需要始终用 $ref 的外部类型
	externalRefTypes := map[string]struct{}{
		"PageData": {},
		"SortData": {},
		"Query":    {},
	}

	switch {
	case strings.HasPrefix(field.Type, "[]"):
		schema["type"] = "array"
		itemType := strings.TrimPrefix(field.Type, "[]")
		if strings.HasPrefix(itemType, "*") {
			itemType = strings.TrimPrefix(itemType, "*")
		}
		if strings.Contains(itemType, ".") {
			parts := strings.Split(itemType, ".")
			itemType = parts[len(parts)-1]
		}
		if _, ok := dtoNames[itemType]; ok {
			schema["items"] = map[string]interface{}{
				"$ref": fmt.Sprintf("#/components/schemas/%s", itemType),
			}
		} else if _, ok := externalRefTypes[itemType]; ok {
			schema["items"] = map[string]interface{}{
				"$ref": fmt.Sprintf("#/components/schemas/%s", itemType),
			}
		} else {
			schema["items"] = map[string]interface{}{
				"type": g.mapGoTypeToOpenAPI(itemType),
			}
		}
	case strings.HasPrefix(field.Type, "*"):
		refType := strings.TrimPrefix(field.Type, "*")
		if strings.Contains(refType, ".") {
			parts := strings.Split(refType, ".")
			refType = parts[len(parts)-1]
		}
		if _, ok := dtoNames[refType]; ok {
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
		} else if _, ok := externalRefTypes[refType]; ok {
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
		} else {
			schema["type"] = g.mapGoTypeToOpenAPI(refType)
		}
	case field.Type == "interface{}":
		schema["type"] = "object"
	case strings.HasPrefix(field.Type, "map["):
		schema["type"] = "object"
		schema["additionalProperties"] = map[string]interface{}{
			"type": "string",
		}
	default:
		if strings.Contains(field.Type, ".") {
			parts := strings.Split(field.Type, ".")
			refType := parts[len(parts)-1]
			if _, ok := dtoNames[refType]; ok {
				schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
			} else if _, ok := externalRefTypes[refType]; ok {
				schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", refType)
			} else {
				schema["type"] = g.mapGoTypeToOpenAPI(refType)
			}
		} else if _, ok := dtoNames[field.Type]; ok {
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", field.Type)
		} else if _, ok := externalRefTypes[field.Type]; ok {
			schema["$ref"] = fmt.Sprintf("#/components/schemas/%s", field.Type)
		} else {
			schema["type"] = g.mapGoTypeToOpenAPI(field.Type)
		}
	}

	if field.Comment != "" {
		schema["description"] = strings.TrimSpace(field.Comment)
	}

	return schema
}
