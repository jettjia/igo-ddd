package global

import (
	"github.com/jinzhu/gorm"

	"github.com/jettjia/go-ddd/infrastructure/config"
	"github.com/jettjia/go-ddd/infrastructure/pkg/log"
)

var (
	Gconfig *config.Config // 全局配置
	GLog    log.Logger     // 全局log
	GDB     *gorm.DB       // 全局 DB
)
