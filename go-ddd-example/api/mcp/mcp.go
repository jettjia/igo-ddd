package mcp

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/jettjia/ddddemo/config"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// MCPService 封装MCP服务

type MCPService struct {
	mcpserver *server.MCPServer
}

// NewMCPService 创建一个新的MCP服务实例
func NewMCPService() *MCPService {
	// 创建MCP服务器
	s := server.NewMCPServer(
		"Agent-Mem MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	return &MCPService{
		mcpserver: s,
	}
}

// RegisterTools 注册MCP工具
func (m *MCPService) RegisterTools() {
	createSysUserTool := mcp.NewTool("create_sys_user",
		mcp.WithDescription("创建用户"),
		mcp.WithString("member_code", mcp.Required(), mcp.Description("会员号")),
		mcp.WithString("phone", mcp.Description("手机号码")),
		mcp.WithString("nick_name", mcp.Description("昵称")),
		mcp.WithString("true_name", mcp.Description("真实姓名")),
		mcp.WithString("level_id", mcp.Description("会员等级id")),
		mcp.WithString("password", mcp.Description("密码")),
		mcp.WithString("admin_level", mcp.Description("1是admin超管")),
		mcp.WithString("created_by", mcp.Description("创建者")),
		mcp.WithString("exp", mcp.Description(`扩展信息，JSON对象，如 {"addr":"...","addr_code":"..."}`)),
	)

	deleteSysUserTool := mcp.NewTool("delete_sys_user",
		mcp.WithDescription("删除用户"),
		mcp.WithString("ulid", mcp.Required(), mcp.Description("用户ulid")),
		mcp.WithString("deleted_by", mcp.Description("删除者")),
	)

	updateSysUserTool := mcp.NewTool("update_sys_user",
		mcp.WithDescription("更新用户"),
		mcp.WithString("ulid", mcp.Required(), mcp.Description("用户ulid")),
		mcp.WithString("member_code", mcp.Description("会员号")),
		mcp.WithString("phone", mcp.Description("手机号码")),
		mcp.WithString("nick_name", mcp.Description("昵称")),
		mcp.WithString("unionid", mcp.Description("微信unionid")),
		mcp.WithString("level_id", mcp.Description("会员等级id")),
		mcp.WithString("updated_by", mcp.Description("修改者")),
	)

	findSysUserByIdTool := mcp.NewTool("find_sys_user_by_id",
		mcp.WithDescription("按ulid查询用户"),
		mcp.WithString("ulid", mcp.Required(), mcp.Description("用户ulid")),
	)

	findSysUserPageTool := mcp.NewTool("find_sys_user_page",
		mcp.WithDescription("分页查询用户"),
		mcp.WithString("query", mcp.Description(`查询条件，JSON数组，如 [{"key":"nick_name","value":"tom","operator":9}]`)),
		mcp.WithString("page_num", mcp.Description("页码，默认1")),
		mcp.WithString("page_size", mcp.Description("每页大小，默认20")),
		mcp.WithString("sort", mcp.Description("排序字段")),
		mcp.WithString("direction", mcp.Description("asc/desc")),
	)

	// 获取处理器实例
	handler := NewMCPHandler()

	// 添加工具处理器
	m.mcpserver.AddTool(createSysUserTool, handler.CreateSysUserHandler)
	m.mcpserver.AddTool(deleteSysUserTool, handler.DeleteSysUserHandler)
	m.mcpserver.AddTool(updateSysUserTool, handler.UpdateSysUserHandler)
	m.mcpserver.AddTool(findSysUserByIdTool, handler.FindSysUserByIdHandler)
	m.mcpserver.AddTool(findSysUserPageTool, handler.FindSysUserPageHandler)

	log.Println("MCP tools registered successfully")
}

// Start 启动MCP服务
func (m *MCPService) Start(ctx context.Context) error {
	// 确保先注册工具
	m.RegisterTools()

	// 获取配置
	cfg := config.NewConfig()
	mcpPort := 7777
	if v, ok := cfg.Third.Extra["mcp_port"]; ok {
		switch t := v.(type) {
		case int:
			mcpPort = t
		case int64:
			mcpPort = int(t)
		case float64:
			mcpPort = int(t)
		case string:
			if i, err := strconv.Atoi(t); err == nil && i > 0 {
				mcpPort = i
			}
		default:
		}
	}
	go func() {
		fmt.Printf("[MCP] Listening and serving MCP on :%d\n", mcpPort)

		// 创建HTTP服务器并直接启动
		httpServer := server.NewStreamableHTTPServer(m.mcpserver)
		if err := httpServer.Start(fmt.Sprintf(":%d", mcpPort)); err != nil {
			fmt.Printf("[MCP] MCP HTTP server error: %v\n", err)
		}
	}()

	// 启动stdio服务器（保持向后兼容）
	if err := server.ServeStdio(m.mcpserver); err != nil {
		return fmt.Errorf("MCP server error: %v", err)
	}

	return nil
}

// GetServer 获取底层MCP服务器实例
func (m *MCPService) GetServer() *server.MCPServer {
	return m.mcpserver
}
