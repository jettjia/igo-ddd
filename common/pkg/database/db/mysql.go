package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	Conf *DBConfig
	Conn *gorm.DB
}

var (
	dbOnce sync.Once
	dbImpl *DB
)

func NewDBClient(cfg *DBConfig) *DB {

	dbOnce.Do(func() {
		dbImpl = &DB{
			Conf: cfg,
		}
		dbImpl.getConn()
	})

	return dbImpl
}

type DBConfig struct {
	Host          string // 服务器地址
	Port          int    // 端口
	DbType        string // db类型
	User          string // 数据库用户名
	Password      string // 数据库密码
	Db            string // 数据名
	DbChar        string // 字符集
	MaxIdleConn   int    // 最大空闲连接
	MaxOpenConn   int    // 最大连接数
	MaxLifetime   int    // 最大生存时间(s)
	LogMode       int    // 是否打印日志
	SlowThreshold int    // 慢sql的起始时间
}

// getConn 链接数据库
func (db *DB) getConn() *gorm.DB {
	var (
		dsn string
		err error
	)
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db.Conf.User, db.Conf.Password, db.Conf.Host, db.Conf.Port, db.Conf.Db, db.Conf.DbChar)

	// gorm logger 配置
	loggerDefault := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Duration(db.Conf.SlowThreshold) * time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.LogLevel(db.Conf.LogMode),                   // Log level
			Colorful:                  true,                                               // 彩色打印
			IgnoreRecordNotFoundError: true,                                               // 关闭 not found错误
		},
	)

	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: loggerDefault, // gorm的log设置
	}

	if db.Conn, err = gorm.Open(mysql.Open(dsn), cfg); err != nil {
		panic(err)
	}

	sqlDB, _ := db.Conn.DB()
	sqlDB.SetMaxIdleConns(db.Conf.MaxOpenConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetMaxOpenConns(db.Conf.MaxOpenConn) //设置数据库连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(db.Conf.MaxLifetime) * time.Second)

	// 使用插件
	if err = db.Conn.Use(&TracePlugin{}); err != nil {
		panic(err)
	}

	return db.Conn
}
