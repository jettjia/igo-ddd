package sys

type SysLog struct {
	Id        uint64 `gorm:"column:id;primary_key;type:bigint(20) unsigned auto_increment;not null;comment:ID;" json:"id"` // ID
	Uuid      string `gorm:"column:uuid;type:varchar(32);comment:uuid;" json:"uuid"`                                       // uuid
	CreatedAt int64  `gorm:"column:created_at;type:bigint(20);comment:创建时间;" json:"created_at"`                            // 创建时间
	CreatedBy string `gorm:"column:created_by;type:varchar(32);comment:创建者;" json:"created_by"`                            // 创建者
	Msg       string `gorm:"column:msg;type:varchar(64); uniqueIndex;comment:msg;" json:"msg"`                             // msg
}
