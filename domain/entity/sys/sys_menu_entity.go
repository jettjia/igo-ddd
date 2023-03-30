package entity

type SysMenu struct {
	Id          uint64 ` json:"id"`           // ID
	CreatedAt   int64  ` json:"created_at"`   // 创建时间
	UpdatedAt   int64  ` json:"updated_at"`   // 修改时间
	DeletedAt   int64  ` json:"deleted_at"`   // 删除时间
	CreatedBy   string ` json:"created_by"`   // 创建者
	UpdatedBy   string ` json:"updated_by"`   // 修改者
	DeletedBy   string ` json:"deleted_by"`   // 删除者
	MenuName    string ` json:"menu_name"`    // menu名称
	Desc        string ` json:"desc"`         // 描述
	Route       string ` json:"route"`        // 菜单路由
	State       uint   ` json:"state"`        // 1显示,2否
	Pid         uint64 ` json:"pid"`          // 父id
	Pname       string ` json:"pname"`        // 父路由名称
	SortOrder   int    ` json:"sort_order"`   // 排序
	BackendType int    ` json:"backend_type"` // 1总后台，2运营后台
}
