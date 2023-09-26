package http

import (
	"fmt"

	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/interface/http/middleware"
	"jettjia/go-ddd-demo-multi-system/interface/http/router"
)

var wg sync.WaitGroup

func InitHttp(server *cmd.Server) {
	wg.Add(2)
	// open api
	go func() {
		defer wg.Done()
		runHttp(server, "public", true)
	}()

	// private api
	go func() {
		defer wg.Done()
		runHttp(server, "private", false)
	}()

	wg.Wait()
}

func runHttp(server *cmd.Server, portType string, jwtEnable bool) {
	var (
		engine *gin.Engine
		port   int
	)
	engineDefault := gin.Default()
	gin.SetMode(server.Cfg.Server.Mode)

	// 配置跨域
	engineDefault.Use(middleware.Cors())
	// auth jwt
	if jwtEnable {
		engineDefault.Use(middleware.TokenAuthorization())
	}
	// 全局recover
	engineDefault.Use(middleware.CatchError())

	// 健康检查
	engineDefault.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engineDefault.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	if portType == "public" {
		engine = router.Routers(engineDefault, server)
		port = server.Cfg.Server.PublicPort
	} else {
		engine = router.RoutersPrivate(engineDefault, server)
		port = server.Cfg.Server.PrivatePort
	}

	err := engine.Run(fmt.Sprintf(":%d", port)) // 启动web
	if err != nil {
		panic(err)
	}
}
