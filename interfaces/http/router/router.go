package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/http/middleware"
	internalRouter "github.com/jettjia/go-ddd-demo/interfaces/http/router/internal"
	sysRouter "github.com/jettjia/go-ddd-demo/interfaces/http/router/sys"
)

func Routers() *gin.Engine {
	gin.SetMode(global.Gconfig.Server.Mode)
	engine := gin.Default()
	// 健康检查
	engine.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engine.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	// 配置跨域
	engine.Use(middleware.Cors())
	// 全局recover
	engine.Use(middleware.Recover())
	// 全局错误
	engine.Use(middleware.ErrorHandler)
	// auth jwt
	engine.Use(middleware.TokenAuthorization())

	// 注册路由
	ApiGroup := engine.Group("/openapi/pc/v1")
	sysRouter.InitSysRouter(ApiGroup) // sys

	return engine
}

func RoutersInternal() *gin.Engine {
	gin.SetMode(global.Gconfig.Server.Mode)
	engine := gin.Default()
	// 健康检查
	engine.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engine.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	// 配置跨域
	engine.Use(middleware.Cors())
	// 全局recover
	engine.Use(middleware.Recover())
	// 全局错误
	engine.Use(middleware.ErrorHandler)

	// 注册路由
	ApiGroup := engine.Group("/api/pc/v1")
	internalRouter.InitInternalRouter(ApiGroup) // sys

	return engine
}
