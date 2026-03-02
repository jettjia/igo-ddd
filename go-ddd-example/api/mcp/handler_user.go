package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"

	dtoUser "github.com/jettjia/ddddemo/application/dto/user"
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"
)

func getStringFromAny(v any, defaultVal string) string {
	if v == nil {
		return defaultVal
	}
	switch t := v.(type) {
	case string:
		if t == "" {
			return defaultVal
		}
		return t
	case []byte:
		if len(t) == 0 {
			return defaultVal
		}
		return string(t)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case json.Number:
		return t.String()
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return defaultVal
		}
		if len(b) == 0 {
			return defaultVal
		}
		return string(b)
	}
}

func getIntFromAny(v any, defaultVal int) int {
	if v == nil {
		return defaultVal
	}
	switch t := v.(type) {
	case int:
		return t
	case int64:
		return int(t)
	case uint:
		return int(t)
	case uint64:
		return int(t)
	case float64:
		return int(t)
	case float32:
		return int(t)
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return defaultVal
		}
		return int(i)
	case string:
		if t == "" {
			return defaultVal
		}
		i, err := strconv.Atoi(t)
		if err != nil {
			return defaultVal
		}
		return i
	default:
		return defaultVal
	}
}

func decodeArgsInto(args map[string]any, out any) error {
	b, err := json.Marshal(args)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, out)
}

func decodeAnyInto(v any, out any) error {
	if v == nil {
		return nil
	}
	switch t := v.(type) {
	case string:
		if t == "" {
			return nil
		}
		return json.Unmarshal([]byte(t), out)
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, out)
	}
}

func (h *MCPHandler) CreateSysUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 进行API Key验证作为双重保障
	if !h.validateApiKeyInHandler(ctx, request.Header) {
		return mcp.NewToolResultError("Unauthorized: Invalid or missing API Key"), nil
	}

	args := request.GetArguments()
	var dtoReq dtoUser.CreateSysUserReq
	if err := decodeArgsInto(args, &dtoReq); err != nil {
		return nil, fmt.Errorf("failed to decode args: %w", err)
	}

	if dtoReq.MemberCode == "" {
		return nil, fmt.Errorf("member_code is required")
	}

	res, err := h.SysUserSrv.CreateSysUser(ctx, &dtoReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create sys_user: %w", err)
	}

	return mcp.NewToolResultText(res.Ulid), nil
}

func (h *MCPHandler) DeleteSysUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 进行API Key验证作为双重保障
	if !h.validateApiKeyInHandler(ctx, request.Header) {
		return mcp.NewToolResultError("Unauthorized: Invalid or missing API Key"), nil
	}

	args := request.GetArguments()
	ulid := getStringFromAny(args["ulid"], "")
	if ulid == "" {
		ulid = request.GetString("ulid", "")
	}
	if ulid == "" {
		return nil, fmt.Errorf("ulid is required")
	}

	err := h.SysUserSrv.DeleteSysUser(ctx, &dtoUser.DelSysUsersReq{
		Ulid: ulid,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to delete sys_user: %w", err)
	}

	return mcp.NewToolResultText("ok"), nil
}

func (h *MCPHandler) UpdateSysUserHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 进行API Key验证作为双重保障
	if !h.validateApiKeyInHandler(ctx, request.Header) {
		return mcp.NewToolResultError("Unauthorized: Invalid or missing API Key"), nil
	}

	args := request.GetArguments()

	var dtoReq dtoUser.UpdateSysUserReq
	if err := decodeArgsInto(args, &dtoReq); err != nil {
		return nil, fmt.Errorf("failed to decode args: %w", err)
	}
	if dtoReq.Ulid == "" {
		dtoReq.Ulid = request.GetString("ulid", "")
	}
	if dtoReq.Ulid == "" {
		return nil, fmt.Errorf("ulid is required")
	}

	err := h.SysUserSrv.UpdateSysUser(ctx, &dtoReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update sys_user: %w", err)
	}

	return mcp.NewToolResultText("ok"), nil
}

func (h *MCPHandler) FindSysUserByIdHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 进行API Key验证作为双重保障
	if !h.validateApiKeyInHandler(ctx, request.Header) {
		return mcp.NewToolResultError("Unauthorized: Invalid or missing API Key"), nil
	}

	args := request.GetArguments()
	ulid := getStringFromAny(args["ulid"], "")
	if ulid == "" {
		ulid = request.GetString("ulid", "")
	}
	if ulid == "" {
		return nil, fmt.Errorf("ulid is required")
	}

	rsp, err := h.SysUserSrv.FindSysUserById(ctx, &dtoUser.FindSysUserByIdReq{Ulid: ulid})
	if err != nil {
		return nil, fmt.Errorf("failed to find sys_user: %w", err)
	}

	b, err := json.Marshal(rsp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(b)), nil
}

func (h *MCPHandler) FindSysUserPageHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if !h.validateApiKeyInHandler(ctx, request.Header) {
		return mcp.NewToolResultError("Unauthorized: Invalid or missing API Key"), nil
	}

	args := request.GetArguments()

	var queries []*builder.Query
	if err := decodeAnyInto(args["query"], &queries); err != nil {
		return nil, fmt.Errorf("invalid query: %w", err)
	}

	pageNum := getIntFromAny(args["page_num"], 1)
	pageSize := getIntFromAny(args["page_size"], 20)
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}

	sort := getStringFromAny(args["sort"], "")
	direction := getStringFromAny(args["direction"], "")

	dtoReq := dtoUser.FindSysUserPageReq{
		Query: queries,
		PageData: &builder.PageData{
			PageNum:  pageNum,
			PageSize: pageSize,
		},
		SortData: &builder.SortData{
			Sort:      sort,
			Direction: direction,
		},
	}

	rsp, err := h.SysUserSrv.FindSysUserPage(ctx, &dtoReq)
	if err != nil {
		return nil, fmt.Errorf("failed to find sys_user page: %w", err)
	}

	b, err := json.Marshal(rsp)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(b)), nil
}
