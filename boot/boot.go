package boot

import (
	"github.com/jettjia/gin-ddd/global"
	"github.com/jettjia/gin-ddd/infrastructure/config"
	"github.com/jettjia/gin-ddd/infrastructure/pkg/database/mysql"
	"github.com/jettjia/gin-ddd/infrastructure/pkg/log"
)

// 初始化
func init() {
	global.Gconfig = config.NewConfig() // 初始化全局配置
	global.GLog = log.NewLogger()       // 初始化日志
	global.GDB = mysql.NewDB()          // 初始化mysql
}
