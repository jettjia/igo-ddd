package mcp

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var mcpservice *MCPService

// InitMCP 初始化并启动MCP服务
func InitMCP() {
	// 创建MCP服务实例
	mcpservice = NewMCPService()

	// 创建一个上下文用于控制服务生命周期
	ctx, cancel := context.WithCancel(context.Background())

	// 在goroutine中启动MCP服务
	go func() {
		// 处理信号
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// 等待信号或上下文取消
		select {
		case sig := <-sigChan:
			log.Printf("Received signal: %v, shutting down MCP service...", sig)
			cancel()
		case <-ctx.Done():
			log.Println("Context canceled, shutting down MCP service...")
		}
	}()

	// 启动服务
	go func() {
		if err := mcpservice.Start(ctx); err != nil {
			log.Printf("MCP service error: %v", err)
			cancel()
		}
	}()

	log.Println("MCP service initialized and started")
}

// GetMCPService 获取MCP服务实例
func GetMCPService() *MCPService {
	return mcpservice
}
