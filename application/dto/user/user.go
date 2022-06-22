package dto

// 输出对象
type (
	SimpleUserInfo struct {
		Id       uint64 `json:"id"`
		Nickname string `json:"nickname"`
	}

	UserLog struct {
		Id  uint64 `json:"log_id"`
		Log string `json:"log"`
	}

	UserInfo struct {
		Id       uint64     `json:"id"`
		Passport string     `json:"passport"`
		Nickname string     `json:"nickname"`
		Logs     []*UserLog `json:"logs"`
	}
)

// 请求对象
type (
	SimpleUserInfoReq struct {
		Id uint64 `json:"id"`
	}
)
