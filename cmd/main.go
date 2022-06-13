package main

import (
	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/interfaces/http/router"
)

func main() {
	router.Routers().Run(global.Gconfig.Server.Address) // 启动web
}
