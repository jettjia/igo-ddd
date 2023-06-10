package http

import (
	"fmt"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/http/router"
)

func InitHttp() {
	// open api
	go func() {
		engine := router.Routers()
		err := engine.Run(fmt.Sprintf(":%d", global.Gconfig.Server.PublicPort))
		if err != nil {
			panic(any(err))
		}
	}()

	// internal api
	go func() {
		engine := router.RoutersInternal()
		err := engine.Run(fmt.Sprintf(":%d", global.Gconfig.Server.PrivatePort))
		if err != nil {
			panic(any(err))
		}
	}()
}
