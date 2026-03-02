package mcp

import (
	"net/http"
	"strings"
)

// validateApiKey 验证API Key是否有效
func validateApiKey(apiKey string, validKeys []string) bool {
	for _, key := range validKeys {
		if key == apiKey {
			return true
		}
	}
	return false
}

// GetApiKeyFromContext 尝试从多种可能的上下文位置提取 API Key。
func GetApiKeyFromContext(header http.Header) string {
	var apiKey string

	if a := header.Get("Authorization"); a != "" {
		if strings.HasPrefix(strings.ToLower(a), "bearer ") {
			apiKey = strings.TrimSpace(a[7:])
		} else {
			apiKey = strings.TrimSpace(a)
		}
	}
	if apiKey == "" {
		if a := header.Get("X-Api-Key"); a != "" {
			apiKey = strings.TrimSpace(a)
		}
		if apiKey == "" {
			if a := header.Get("Api-Key"); a != "" {
				apiKey = strings.TrimSpace(a)
			}
		}
	}

	return apiKey
}
