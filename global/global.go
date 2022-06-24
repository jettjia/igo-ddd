package global

import (
	"github.com/jinzhu/gorm"
	"github.com/nsqio/go-nsq"

	"github.com/jett/gin-ddd/infrastructure/config"
	"github.com/jett/gin-ddd/infrastructure/pkg/log"
)

var (
	Gconfig      *config.Config // 全局配置
	GLog         log.Logger     // 全局log
	GDB          *gorm.DB       // 全局 DB
	GNsqProducer *nsq.Producer  // 全局 nsq producer
)
