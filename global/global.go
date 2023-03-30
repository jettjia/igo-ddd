package global

import (
	"gorm.io/gorm"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/jwt"
	"github.com/jettjia/go-ddd-demo/infrastructure/config"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/log"
)

var (
	Gconfig       *config.Config      // 全局配置
	GLog          log.Logger          // 全局log
	GDB           *gorm.DB            // 全局 DB
	GCustomerInfo entity.CustomerInfo // 登录的全局信息
)
