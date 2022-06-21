package util

import "os"

// Getwd 获得项目的根路径
func Getwd() string {
	dir, _ := os.Getwd()
	return dir
}
