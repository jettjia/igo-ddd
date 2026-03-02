package mcp

import (
	"context"
	"net/http"

	"github.com/jettjia/ddddemo/config"

	serviceUser "github.com/jettjia/ddddemo/application/service/user"
)

// MCPHandler 处理MCP工具调用的处理器
type MCPHandler struct {
	SysUserSrv *serviceUser.SysUserService
}

// NewMCPHandler 创建一个新的MCP处理器实例
func NewMCPHandler() *MCPHandler {
	return &MCPHandler{
		SysUserSrv: serviceUser.NewSysUserService(),
	}
}

// validateApiKeyInHandler 在处理器中验证API Key
func (h *MCPHandler) validateApiKeyInHandler(ctx context.Context, header http.Header) bool {
	// 获取配置
	cfg := config.NewConfig()

	// 检查是否启用认证
	enableAuth, ok := cfg.Third.Extra["mcp_enable_auth"].(bool)
	if !ok || !enableAuth {
		// 如果未启用认证，直接返回true
		return true
	}

	apiKey := GetApiKeyFromContext(header)

	// 从配置中获取有效的API Key列表
	apiKeys, ok := cfg.Third.Extra["mcp_api_keys"].([]interface{})
	if !ok {
		return false
	}

	// 将interface{}切片转换为string切片
	validKeys := make([]string, 0, len(apiKeys))
	for _, key := range apiKeys {
		if keyStr, ok := key.(string); ok {
			validKeys = append(validKeys, keyStr)
		}
	}

	// 验证API Key
	return validateApiKey(apiKey, validKeys)
}
