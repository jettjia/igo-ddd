package http

import (
	"github.com/gin-gonic/gin"

	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/interfaces/http/handler/user"
	"github.com/jett/gin-ddd/interfaces/http/middleware"
)

type NewHttpApp struct {
	URESTHandler user.RESTHandler // 注入用户handler
}

func (app *NewHttpApp) Start() {
	gin.SetMode(global.Gconfig.Server.Mode)
	engine := gin.New()

	// 中间件
	engine.Use(middleware.Cors())
	engine.Use(middleware.ErrorHandler)

	// 注册用户API
	app.URESTHandler.RegisterAPI(engine)

	if err := engine.Run(global.Gconfig.Server.Address); err != nil {
		panic(err)
	}
}

func init() {
	server := NewHttpApp{
		URESTHandler: user.NewRESTHandler(), // 注册用户api
	}
	server.Start()
}
