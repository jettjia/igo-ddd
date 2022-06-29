package http

import (
	"github.com/jettjia/go-ddd/cmd"
	"github.com/jettjia/go-ddd/global"
	"github.com/jettjia/go-ddd/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	go func() {
		router.Routers(app).Run(global.Gconfig.Server.Address) // 启动web
	}()
}
