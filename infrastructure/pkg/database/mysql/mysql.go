package mysql

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jettjia/go-ddd/global"
)

func NewDB() *gorm.DB {
	return initDB()
}

func initDB() *gorm.DB {
	c := global.Gconfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.DbHost, c.DbPort, c.DbName)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(true)       //打印sql语句

	db.SetLogger(&MyLogger{})

	db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.DB().SetConnMaxLifetime(30 * time.Second)

	return db
}

type MyLogger struct {
}

func (logger *MyLogger) Print(values ...interface{}) {
	var (
		level  = values[0]
		source = values[1].(string)
		doTime = values[2]
		sql    = values[3].(string)
	)

	if level == "sql" {
		logStr := fmt.Sprintf("%s", doTime) + " " + source + " " + sql
		fmt.Println(logStr)
	}
}
