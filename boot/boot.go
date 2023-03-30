package boot

import (
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/config"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/database/mysql"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/log"
)

func InitServer(env string) {
	global.Gconfig = config.NewConfig(env) // 初始化全局配置
	global.GLog = log.NewLogger(env)       // 初始化日志
	global.GDB = mysql.NewDB()             // 初始化mysql
}
