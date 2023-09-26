package dto

import "jettjia/go-ddd-demo-multi-common/types"

// 请求对象
type (
	// CreateSysMenuReq 创建SysMenu 请求对象
	CreateSysMenuReq struct {
		CreatedBy   string `json:"created_by"`                                                     // 创建者
		MenuName    string `validate:"required,min=1,max=4" err_info:"长度在1-4个字符" json:"menu_name"` // menu名称 【校验说明：肯定会校验，长度在1-4】
		Desc        string `validate:"omitempty,min=1,max=32" json:"desc"`                         // 描述 【校验说明：有就会校验，长度在1-32】
		Route       string `validate:"required,checkSpecialChar" json:"route"`                     // 菜单路由 【校验说明：自定义校验规则】
		State       uint   `validate:"required,oneof=1 2" json:"state"`                            // 1显示,2否 【校验说明：肯定校验，并且值只能是1或者2】
		Pid         string `json:"pid"`                                                            // 父id
		Pname       string `json:"pname"`                                                          // 父路由名称
		SortOrder   int    `json:"sort_order"`                                                     // 排序
		BackendType int    `json:"backend_type"`                                                   // 1总后台，2运营后台
	}

	// DelSysMenusReq 删除 请求对象
	DelSysMenusReq struct {
		Ulid      string `validate:"required" uri:"ulid" json:"ulid"` // ID
		DeletedBy string `json:"deleted_by"`                          // 删除者
	}

	// UpdateSysMenuReq 修改SysMenu 请求对象
	UpdateSysMenuReq struct {
		Ulid        string `validate:"required" uri:"ulid" json:"ulid"` // ID
		UpdatedBy   string `json:"updated_by"`                          // 修改者
		MenuName    string `json:"menu_name"`                           // menu名称
		Desc        string `json:"desc"`                                // 描述
		Route       string `json:"route"`                               // 菜单路由
		State       uint   `json:"state"`                               // 1显示,2否
		Pid         string `json:"pid"`                                 // 父id
		Pname       string `json:"pname"`                               // 父路由名称
		SortOrder   int    `json:"sort_order"`                          // 排序
		BackendType int    `json:"backend_type"`                        // 1总后台，2运营后台
	}

	// FindSysMenuByIdReq 查询 请求对象
	FindSysMenuByIdReq struct {
		Ulid string ` validate:"required" uri:"ulid" json:"ulid"` // ID
	}

	// FindSysMenuByQueryReq 查询 请求对象
	FindSysMenuByQueryReq struct {
		Query []*types.Query `json:"query"`
	}

	// FindSysMenuAllReq 查询 请求对象
	FindSysMenuAllReq struct {
		Query []*types.Query `json:"query"`
	}

	// FindSysMenuPageReq 分页查询 请求对象
	FindSysMenuPageReq struct {
		Query    []*types.Query  `json:"query"`
		PageData *types.PageData `json:"page_data"`
		SortData *types.SortData `json:"sort_data"`
	}
)

// 输出对象
type (
	// CreateSysMenuRsp 创建SysMenu 返回对象
	CreateSysMenuRsp struct {
		Ulid string `json:"ulid"` // ID
	}

	// FindSysMenuPageRsp 列表查询 返回对象
	FindSysMenuPageRsp struct {
		Entries  []*FindSysMenuRsp `json:"entries"`
		PageData *types.PageData   `json:"page_data"`
	}

	// FindSysMenuRsp 查询SysMenu 返回对象
	FindSysMenuRsp struct {
		Ulid        string `json:"ulid"`         // ID
		CreatedAt   int64  `json:"created_at"`   // 创建时间
		UpdatedAt   int64  `json:"updated_at"`   // 修改时间
		CreatedBy   string `json:"created_by"`   // 创建者
		UpdatedBy   string `json:"updated_by"`   // 修改者
		MenuName    string `json:"menu_name"`    // menu名称
		Desc        string `json:"desc"`         // 描述
		Route       string `json:"route"`        // 菜单路由
		State       uint   `json:"state"`        // 1显示,2否
		Pid         string `json:"pid"`          // 父id
		Pname       string `json:"pname"`        // 父路由名称
		SortOrder   int    `json:"sort_order"`   // 排序
		BackendType int    `json:"backend_type"` // 1总后台，2运营后台
	}
)
