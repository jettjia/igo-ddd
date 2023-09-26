package router

import (
	"github.com/gin-gonic/gin"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/interface/http/handler/private"
	"jettjia/go-ddd-demo-multi-system/interface/http/handler/public"
	privateRouter "jettjia/go-ddd-demo-multi-system/interface/http/router/private"
	publicRouter "jettjia/go-ddd-demo-multi-system/interface/http/router/public"
)

func Routers(engine *gin.Engine, server *cmd.Server) *gin.Engine {
	// 注册路由
	ApiGroup := engine.Group("/api/pc/v1")
	hand := public.NewHandler(server.Sys)

	publicRouter.SetPublicRouter(ApiGroup, hand) // sys

	return engine
}

func RoutersPrivate(engine *gin.Engine, server *cmd.Server) *gin.Engine {
	// 注册路由
	hand := private.NewPrivateHandler()

	ApiGroup := engine.Group("/private/pc/v1")
	privateRouter.SetPrivateRouter(ApiGroup, hand) // sys

	return engine
}
