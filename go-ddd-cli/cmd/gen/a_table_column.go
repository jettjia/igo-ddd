package gen

import (
	"fmt"
	"log"
	"os"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GetTableCol(table string) []TableColumn {
	var tc []TableColumn
	var query string

	// 获取数据库类型
	dialectName := DB.Dialector.Name()

	if dialectName == "postgres" {
		query = `SELECT
			column_name as "Field",
			CASE
				WHEN udt_name = 'varchar' THEN data_type || '(' || character_maximum_length || ')'
				WHEN udt_name = 'numeric' THEN data_type || '(' || numeric_precision || ',' || numeric_scale || ')'
				ELSE data_type
			END as "Type",
			CASE WHEN is_nullable = 'YES' THEN 'YES' ELSE 'NO' END as "Null",
			COALESCE(column_default, '') as "Default",
			COALESCE(col_description((table_schema || '.' || table_name)::regclass::oid, ordinal_position), '') as "Comment",
			CASE WHEN column_default LIKE 'nextval%' THEN 'auto_increment' ELSE '' END as "Extra"
		FROM information_schema.columns
		WHERE table_name = '` + table + `'
		ORDER BY ordinal_position`
	} else {
		query = "SHOW FULL COLUMNS FROM `" + table + "`"
	}

	err := DB.Raw(query).Scan(&tc).Error
	if err != nil {
		panic(any(err))
	}

	fmt.Println("tc", tc)

	return tc
}

func InitDB(host, port, user, password, dbName, dbType string) {
	var db *gorm.DB
	var err error

	// sql_gorm logger 配置
	loggerDefault := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.LogLevel(4), // Log level
			Colorful:                  true,               // 彩色打印
			IgnoreRecordNotFoundError: true,               // 关闭 not found错误
		},
	)

	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: loggerDefault, // gorm的log设置
	}

	if dbType == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
		mlog.Print("db conn: ", dsn)
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			panic(any(err))
		}
	} else {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbName, "utf8mb4")
		mlog.Print("db conn: ", dsn)
		db, err = gorm.Open(mysql.Open(dsn), cfg)
		if err != nil {
			panic(any(err))
		}
	}

	DB = db
}
