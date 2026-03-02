package router

import (
	"github.com/gin-gonic/gin"

	publicRouter "github.com/jettjia/ddddemo/api/http/router/public"
)

func Routers(engine *gin.Engine) *gin.Engine {
	// 注册路由
	ApiGroup := engine.Group("/api/xtext/x-data/v1")
	publicRouter.SetPublicRouter(ApiGroup) // router
	return engine
}
