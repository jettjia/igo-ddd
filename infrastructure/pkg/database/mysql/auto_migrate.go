package mysql

import (
	"gorm.io/gorm"

	sysPo "github.com/jettjia/go-ddd-demo/infrastructure/repository/po/sys"
)

func MysqlAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&sysPo.SysMenu{},
	)

	if err != nil {
		panic(any(err))
	}
}
