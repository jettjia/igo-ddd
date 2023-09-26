package data

import (
	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Data struct {
	Mysql      *gorm.DB
	RedisCli   redis.UniversalClient
	RocksCache *rockscache.Client
}

func NewData(mysqlDB *gorm.DB, redisCli redis.UniversalClient, rocksCli *rockscache.Client) (*Data, error) {
	return &Data{
		Mysql:      mysqlDB,
		RedisCli:   redisCli,
		RocksCache: rocksCli,
	}, nil
}
