package http

import (
	"fmt"
	"github.com/jettjia/go-ddd-demo/boot"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/http/router"
)

func InitHttp(app *boot.App) {
	// open api
	go func() {
		engine := router.Routers(app)
		err := engine.Run(fmt.Sprintf(":%d", global.Gconfig.Server.PublicPort))
		if err != nil {
			panic(any(err))
		}
	}()

	// internal api
	go func() {
		engine := router.RoutersInternal(app)
		err := engine.Run(fmt.Sprintf(":%d", global.Gconfig.Server.PrivatePort))
		if err != nil {
			panic(any(err))
		}
	}()
}
