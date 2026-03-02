package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"tools/apidoc/generator"
	"tools/apidoc/parser"
)

func findProjectRoot() (string, error) {
	// 从当前目录开始向上查找，直到找到包含 go.mod 的目录
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// 检查当前目录是否包含 go.mod
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// 再向上一级，因为我们的工具在 tools/apidoc 目录下
			return filepath.Dir(filepath.Dir(dir)), nil
		}

		// 获取父目录
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("could not find project root (no go.mod found)")
		}
		dir = parent
	}
}

func main() {
	// 找到项目根目录
	projectRoot, err := findProjectRoot()
	if err != nil {
		log.Fatalf("Failed to find project root: %v", err)
	}

	// 配置扫描路径（使用绝对路径）
	config := &parser.Config{
		RouterPath:  filepath.Join(projectRoot, "interface/http/router"),
		HandlerPath: filepath.Join(projectRoot, "interface/http/handler"),
		DtoPath:     filepath.Join(projectRoot, "application/dto"),
	}

	// 解析路由
	routerParser := parser.NewRouterParser(config)
	routes, err := routerParser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse routers: %v", err)
	}

	fmt.Printf("=== 解析到的路由信息 ===\n")
	for _, route := range routes {
		fmt.Printf("Path: %s, Method: %s, Handler: %s.%s\n", route.Path, route.Method, route.HandlerPkg, route.Handler)
	}

	// 解析 handlers
	handlerParser := parser.NewHandlerParser(config)
	handlers, err := handlerParser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse handlers: %v", err)
	}

	fmt.Printf("\n=== 解析到的 Handler 信息 ===\n")
	for _, handler := range handlers {
		fmt.Printf("Handler: %s, Request: %v, Response: %v\n", handler.Name, handler.Request, handler.Response)
	}

	// 解析 DTOs
	dtoParser := parser.NewDtoParser(config)
	dtos, err := dtoParser.Parse()
	if err != nil {
		log.Fatalf("Failed to parse DTOs: %v", err)
	}

	fmt.Printf("\n=== 解析到的 DTO 信息 ===\n")
	for _, dto := range dtos {
		fmt.Printf("DTO: %s (包: %s), 字段数: %d\n", dto.Name, dto.Package, len(dto.Fields))
	}

	// 生成 OpenAPI 文档
	gen := generator.NewOpenAPIGenerator(&generator.Config{
		Title:       "X-Data API",
		Version:     "1.0.0",
		Description: "X-Data REST API documentation",
	})

	doc := gen.Generate(routes, handlers, dtos)

	// 手动补充 builder 包常用类型 schema
	schemas := doc["components"].(map[string]interface{})["schemas"].(map[string]interface{})

	schemas["PageData"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"page_num": map[string]interface{}{
				"type":        "integer",
				"description": "page number",
			},
			"page_size": map[string]interface{}{
				"type":        "integer",
				"description": "the number of rows displayed per page",
			},
			"total_number": map[string]interface{}{
				"type":        "integer",
				"description": "how many in total",
			},
			"total_page": map[string]interface{}{
				"type":        "integer",
				"description": "how many pages in total",
			},
		},
	}

	schemas["SortData"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"sort": map[string]interface{}{
				"type":        "string",
				"description": "sort fields",
			},
			"direction": map[string]interface{}{
				"type":        "string",
				"description": "asc;desc",
			},
		},
	}

	schemas["Query"] = map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"key": map[string]interface{}{
				"type":        "string",
				"description": "the key to search for a keyword",
			},
			"value": map[string]interface{}{
				"type":        "string",
				"description": "search for the value of the keyword",
			},
			"operator": map[string]interface{}{
				"type":        "integer",
				"description": "judging conditions",
			},
		},
	}

	// 输出到文件（在项目根目录下创建 docs 目录）
	outputPath := filepath.Join(projectRoot, "docs/api/swagger.json")
	err = os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		log.Fatalf("Failed to create docs directory: %v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create swagger.json: %v", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(doc); err != nil {
		log.Fatalf("Failed to write swagger.json: %v", err)
	}

	fmt.Printf("\nOpenAPI documentation generated at %s\n", outputPath)
}
