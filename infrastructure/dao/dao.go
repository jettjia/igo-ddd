package dao

import (
	"github.com/jinzhu/gorm"
)

type dao struct {
	mysql *gorm.DB
}

func newDao(mysql *gorm.DB) *dao {
	return &dao{mysql: mysql}
}
