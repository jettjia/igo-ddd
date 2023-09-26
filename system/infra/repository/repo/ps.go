package repo

import (
	"github.com/google/wire"

	"jettjia/go-ddd-demo-multi-common/pkg/data"
)

var RepoProviderSet = wire.NewSet(data.NewData, NewDB, NewRedis, NewRocksCache, data.NewTransaction, NewSysMenuImpl, NewSysLogImpl)
