package entity

type SysLog struct {
	Id        uint64 ` json:"id"`         // ID
	Uuid      string ` json:"uuid"`       // uuid
	CreatedAt int64  ` json:"created_at"` // 创建时间
	CreatedBy string ` json:"created_by"` // 创建者
	Msg       string ` json:"msg"`        // msg
}
