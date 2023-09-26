package repo

import (
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"jettjia/go-ddd-demo-multi-common/pkg/conf"
	"jettjia/go-ddd-demo-multi-common/pkg/database/db"
	redis2 "jettjia/go-ddd-demo-multi-common/pkg/redis"
)

// NewDB gorm Connecting to a Database
func NewDB(conf *conf.Config) *gorm.DB {

	var cfg db.DBConfig
	cfg.Host = conf.Mysql.DbHost
	cfg.Port = conf.Mysql.DbPort
	cfg.User = conf.Mysql.Username
	cfg.Password = conf.Mysql.Password
	cfg.Db = conf.Mysql.DbName
	cfg.DbChar = conf.Mysql.Charset
	cfg.MaxIdleConn = conf.Mysql.MaxIdleConn
	cfg.MaxOpenConn = conf.Mysql.MaxOpenConn
	cfg.MaxLifetime = conf.Mysql.ConnMaxLifetime
	cfg.LogMode = conf.Mysql.LogMode
	cfg.SlowThreshold = conf.Mysql.SlowThreshold

	return db.NewDBClient(&cfg).Conn
}

func NewRedis(conf *conf.Config) redis.UniversalClient {
	cfg := &redis2.RedisConfig{
		RedisType: conf.Redis.RedisType,
		Addrs: []string{
			conf.Redis.Addr,
		},
		Password: conf.Redis.Password,
		PoolSize: conf.Redis.PoolSize,
	}
	return redis2.NewRedisClient(cfg).Conn
}

func NewRocksCache(rdb redis.UniversalClient) *rockscache.Client {
	var dc = rockscache.NewClient(rdb, rockscache.NewDefaultOptions())
	// 常用参数设置
	// 1、强一致性(默认关闭强一致性，如果开启的话会影响性能)
	dc.Options.StrongConsistency = false

	// 2、redis出现问题需要缓存降级时设置为true
	dc.Options.DisableCacheRead = false   // 关闭缓存读，默认false；如果打开，那么Fetch就不从缓存读取数据，而是直接调用fn获取数据
	dc.Options.DisableCacheDelete = false // 关闭缓存删除，默认false；如果打开，那么TagAsDeleted就什么操作都不做，直接返回

	// 3、其他设置
	// 标记删除的延迟时间，默认10秒，设置为13秒表示：被删除的key在13秒后才从redis中彻底清除
	dc.Options.Delay = time.Second * time.Duration(13)
	// 防穿透: 若fn返回空字符串，空结果在缓存中的缓存时间，默认60秒
	dc.Options.EmptyExpire = time.Second * time.Duration(120)
	// 防雪崩: 默认0.1,当前设置为0.1的话，如果设定为600的过期时间，那么过期时间会被设定为540s - 600s中间的一个随机数，避免数据出现同时到期
	dc.Options.RandomExpireAdjustment = 0.1 // 设置为默认或不设置就行

	// 锁相关参数，这里配置的默认值，没有特殊情况建议默认
	// 是更新缓存时分配的锁的过期时间。默认为 3s。设置为下级计算数据时间的最大值。
	dc.Options.LockExpire = time.Second * time.Duration(3)
	// 锁失败后的重试等待时间 100ms
	dc.Options.LockSleep = time.Millisecond * time.Duration(100)

	return dc
}
