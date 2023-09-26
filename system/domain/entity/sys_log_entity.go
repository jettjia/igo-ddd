package entity

type SysLog struct {
	Ulid      string `json:"ulid"`       // ulid
	CreatedAt int64  `json:"created_at"` // 创建时间
	CreatedBy string `json:"created_by"` // 创建者
	Msg       string `json:"msg"`        // msg
}
