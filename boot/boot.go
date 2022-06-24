package boot

import (
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/infrastructure/config"
	"github.com/jett/gin-ddd/infrastructure/pkg/database/mysql"
	"github.com/jett/gin-ddd/infrastructure/pkg/log"
	"github.com/jett/gin-ddd/infrastructure/pkg/mq/nsq"
)

// 初始化
func init() {
	global.Gconfig = config.NewConfig()        // 初始化全局配置
	global.GLog = log.NewLogger()              // 初始化日志
	global.GDB = mysql.NewDB()                 // 初始化mysql
	global.GNsqProducer = nsq.NewNsqProducer() // 初始化 nsq producer
}
