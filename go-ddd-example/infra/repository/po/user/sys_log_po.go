package user

import (
	"github.com/jettjia/igo-pkg/pkg/util"
	"gorm.io/gorm"
)

type SysLog struct {
	Ulid      string `gorm:"column:ulid;primary_key;type:varchar(128)not null;comment:ulid;" json:"ulid"`         // ulid
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli;type:bigint;comment:创建时间;" json:"created_at"`  // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime:milli;type:bigint;comment:修改时间;" json:"updated_at"`  // 修改时间
	DeletedAt int64  `gorm:"column:deleted_at;autoDeletedTime:milli;type:bigint;comment:删除时间;" json:"deleted_at"` // 删除时间
	CreatedBy string `gorm:"column:created_by;type:varchar(32);comment:创建者;" json:"created_by"`                   // 创建者
	UpdatedBy string `gorm:"column:updated_by;type:varchar(32);comment:修改者;" json:"updated_by"`                   // 修改者
	DeletedBy string `gorm:"column:deleted_by;type:varchar(32);comment:删除者;" json:"deleted_by"`                   // 删除者
	Msg       string `gorm:"column:msg;type:varchar(255); uniqueIndex;comment:msg;" json:"msg"`                   // msg
}

func (po *SysLog) BeforeCreate(tx *gorm.DB) (err error) {
	po.Ulid = util.Ulid()
	return
}

func (po *SysLog) TableName() string {
	return "sys_log"
}
