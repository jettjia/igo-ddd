package idata

import (
	"fmt"
	"sync"

	"github.com/jettjia/ddddemo/config"
	"github.com/jettjia/igo-pkg/pkg/database/cache"
)

var (
	cachedKVCache cache.KVCache
	kvCacheOnce   sync.Once
)

func NewKVCacheClient() cache.KVCache {
	kvCacheOnce.Do(func() {
		conf := config.NewConfig()
		var err error
		cachedKVCache, err = cache.NewKVCache("pgsql",
			cache.WithAddr(fmt.Sprintf("%s:%d", conf.DB.DbHost, conf.DB.DbPort)),
			cache.WithPassword(conf.DB.Password))
		if err != nil {
			panic(err)
		}
	})

	return cachedKVCache
}
