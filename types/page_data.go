package types

// PageData 分页
type PageData struct {
	PageNum     int   `json:"page_num"`     // 页码
	PageSize    int   `json:"page_size"`    // 每页显示行数
	TotalNumber int64 `json:"total_number"` // 共多少条
	TotalPage   int64 `json:"total_page"`   // 共多少页
}

// SortData 排序
type SortData struct {
	Sort      string `json:"sort"`      // 排序字段
	Direction string `json:"direction"` // asc：升序;desc：降序
}
