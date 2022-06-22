package http

import (
	service "github.com/jett/gin-ddd/application/service/user"
	"github.com/jett/gin-ddd/global"
	"github.com/jett/gin-ddd/interfaces/http/registersrv"
	"github.com/jett/gin-ddd/interfaces/http/router"
)

func init() {
	server := registersrv.Registersrv{
		UserSrv: service.NewUserService(),
	}

	go func() {
		router.Routers(&server).Run(global.Gconfig.Server.Address) // 启动web
	}()
}
