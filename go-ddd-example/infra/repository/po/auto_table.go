package po

import (
	"github.com/jettjia/ddddemo/config"
	"github.com/jettjia/ddddemo/infra/repository/po/user"
	"github.com/jettjia/igo-pkg/pkg/database/db"
)

// AutoTable auto create table
func AutoTable() (err error) {
	conf := config.NewConfig()
	dbCli := db.NewDBClient(conf).Conn

	err = dbCli.AutoMigrate(
		user.SysUser{},
		user.SysLog{},
	)

	return
}
