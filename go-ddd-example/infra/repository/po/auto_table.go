package po

import (
	"github.com/jettjia/igo-pkg/pkg/database/mysql"

	"github.com/jettjia/ddddemo/config"
	"github.com/jettjia/ddddemo/infra/repository/po/user"
)

// AutoTable auto create table
func AutoTable() (err error) {
	conf := config.NewConfig()
	dbCli := mysql.NewDBClient(conf).Conn

	err = dbCli.AutoMigrate(
		user.SysUser{},
		user.SysLog{},
	)

	return
}
