package http

import (
	"github.com/jettjia/go-ddd-demo/cmd"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	// open api
	go func() {
		engine := router.Routers(app)
		err := engine.Run(global.Gconfig.Server.PublicPort) // 启动web
		if err != nil {
			panic(any(err))
		}
	}()

	// internal api
	go func() {
		engine := router.RoutersInternal(app)
		err := engine.Run(global.Gconfig.Server.PrivatePort) // 启动web
		if err != nil {
			panic(any(err))
		}
	}()
}
