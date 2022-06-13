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
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(middleware.Cors()) // 跨域

	// 注册API
	app.URESTHandler.RegisterAPI(engine)

	if err := engine.Run(global.Gconfig.Server.Address); err != nil {
		panic(err)
	}
}
