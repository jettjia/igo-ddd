package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Conn redis.UniversalClient
}

type RedisConfig struct {
	RedisType  string // redis使用模式:alone, sentinel,cluster
	Addrs      []string
	Password   string
	MasterName string
	PoolSize   int
}

var (
	once   sync.Once
	rdConn *RedisDB
)

// NewRedisClient 获取redis的链接
func NewRedisClient(cfg *RedisConfig) *RedisDB {

	once.Do(func() {
		rdConn = &RedisDB{}

		rdConn.Conn = getConn(cfg)
	})

	return rdConn
}

func getConn(cfg *RedisConfig) redis.UniversalClient {
	var (
		rdb redis.UniversalClient
		typ string
	)

	typ = cfg.RedisType

	switch typ {
	case "alone":
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    cfg.Addrs,
			Password: cfg.Password,
			PoolSize: cfg.PoolSize,
		})
	case "sentinel":
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:      cfg.Addrs,
			MasterName: cfg.MasterName,
			Password:   cfg.Password,
			// To route commands by latency or randomly, enable one of the following.
			RouteByLatency: true,
			RouteRandomly:  true,
			PoolSize:       cfg.PoolSize,
		})
	case "cluster":
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    cfg.Addrs,
			Password: cfg.Password,
			// To route commands by latency or randomly, enable one of the following.
			RouteByLatency: true,
			RouteRandomly:  true,
			PoolSize:       cfg.PoolSize,
		})
	default:
		panic("redis link type error, Link type must be:alone,sentinel,cluster")
	}

	return rdb
}
