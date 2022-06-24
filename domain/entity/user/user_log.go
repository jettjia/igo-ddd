package entity

type UserLog struct {
	ID     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	Log    string `json:"log"`
}
