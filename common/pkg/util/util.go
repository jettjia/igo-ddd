package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
)

// CeilPageNum 分页数目计算
func CeilPageNum(total int64, pageSize int) int64 {
	return int64(int(math.Ceil(float64(total) / float64(pageSize))))
}

// PrintJson 打印Json
func PrintJson(args interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", args)
	}
	return out.String()
}
