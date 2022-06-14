package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jett/gin-ddd/interfaces/http/handler/health"

	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/interfaces/http/handler/user"
	"github.com/jett/gin-ddd/interfaces/http/middleware"
)

type NewHttpApp struct {
	HealthHandler health.RESTHandler // 健康检测
	UserHandler   user.RESTHandler   // 用户模块

}

func (app *NewHttpApp) Start() {
	go func() {
		gin.SetMode(global.Gconfig.Server.Mode)
		engine := gin.New()

		// 中间件
		engine.Use(middleware.Cors())
		engine.Use(middleware.ErrorHandler)

		// 注册用户API
		app.UserHandler.RegisterAPI(engine)
		app.HealthHandler.RegisterAPI(engine)

		logConf := gin.LoggerConfig{
			SkipPaths: []string{"/health/ready", "/health/alive"},
		}
		engine.Use(gin.LoggerWithConfig(logConf))
		if err := engine.Run(global.Gconfig.Server.Address); err != nil {
			panic(err)
		}
	}()
}

func init() {
	server := NewHttpApp{
		HealthHandler: health.NewRESTHandler(), // 健康检测
		UserHandler:   user.NewRESTHandler(),   // 用户模块
	}
	server.Start()
}
