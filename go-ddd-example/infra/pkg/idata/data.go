package idata

import (
	"github.com/jettjia/ddddemo/config"
	"github.com/jettjia/igo-pkg/pkg/data"
	"github.com/jettjia/igo-pkg/pkg/database/db"
	"github.com/jettjia/igo-pkg/pkg/database/dbresolver"
)

// 自定义的引入，依赖的存储包。推荐
func NewDataOptionCli() (dataCli *data.Data) {
	conf := config.NewConfig()

	dbCli := db.NewDBClient(conf).Conn

	// 创建 DBManagerDynamic 实例
	dbManagerDynamic := dbresolver.NewDBManagerDynamic(dbCli, conf)

	dataCli = data.NewDataOption(
		data.WithMysql(dbCli),
		data.WithDBManagerDynamic(dbManagerDynamic),
	)

	return dataCli
}
