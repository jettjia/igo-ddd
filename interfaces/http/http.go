package http

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/jett/gin-ddd/cmd"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/interfaces/http/router"
)

func InitHttp(app *cmd.App) {
	go func() {
		router.Routers(app).Run(global.Gconfig.Server.Address) // 启动web
	}()

	{
		//接收终止信号
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
	}
}
