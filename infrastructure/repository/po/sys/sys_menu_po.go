package po

import (
	"gorm.io/plugin/soft_delete"
)

type SysMenu struct {
	Id          uint64                `gorm:"column:id;primary_key;type:bigint(20) unsigned auto_increment;;comment:ID;" json:"id"`     // ID
	Uuid        string                `gorm:"column:uuid;type:varchar(32);comment:uuid;" json:"uuid"`                                   // uuid
	CreatedAt   int64                 `gorm:"column:created_at;type:bigint(20);;comment:创建时间;" json:"created_at"`                       // 创建时间
	UpdatedAt   int64                 `gorm:"column:updated_at;type:bigint(20);;comment:修改时间;" json:"updated_at"`                       // 修改时间
	DeletedAt   soft_delete.DeletedAt `gorm:"column:deleted_at;comment:删除时间;" json:"deleted_at"`                                        // 删除时间
	CreatedBy   string                `gorm:"column:created_by;type:varchar(32);;comment:创建者;" json:"created_by"`                       // 创建者
	UpdatedBy   string                `gorm:"column:updated_by;type:varchar(32);;comment:修改者;" json:"updated_by"`                       // 修改者
	DeletedBy   string                `gorm:"column:deleted_by;type:varchar(32);;comment:删除者;" json:"deleted_by"`                       // 删除者
	MenuName    string                `gorm:"uniqueIndex:uqx_name;column:menu_name;type:varchar(32);;comment:menu名称;" json:"menu_name"` // menu名称
	Desc        string                `gorm:"column:desc;type:varchar(255);;comment:描述;" json:"desc"`                                   // 描述
	Route       string                `gorm:"column:route;type:varchar(128);;comment:菜单路由;" json:"route"`                               // 菜单路由
	State       uint                  `gorm:"column:state;type:int(1) unsigned;;default: 1;comment:1显示,2否;" json:"state"`               // 1显示,2否
	Pid         uint64                `gorm:"column:pid;type:bigint(20) unsigned;;default: 0;comment:父id;" json:"pid"`                  // 父id
	Pname       string                `gorm:"column:pname;type:varchar(32);;comment:父路由名称;" json:"pname"`                               // 父路由名称
	SortOrder   int                   `gorm:"column:sort_order;type:smallint(6);;default: 50;comment:排序;" json:"sort_order"`            // 排序
	BackendType int                   `gorm:"column:backend_type;type:tinyint(1);;default: 1;comment:1总后台，2运营后台;" json:"backend_type"`  // 1总后台，2运营后台
}
