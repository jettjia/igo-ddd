package http

import (
	"github.com/jettjia/gin-ddd/cmd"
	"github.com/jettjia/gin-ddd/global"
	"github.com/jettjia/gin-ddd/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	go func() {
		router.Routers(app).Run(global.Gconfig.Server.Address) // 启动web
	}()
}
