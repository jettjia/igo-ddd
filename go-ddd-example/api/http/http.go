package http

import (
	"context"
	"fmt"
	"net/http"
	"os"

	tel "github.com/jettjia/igo-pkg/pkg/otel"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/jettjia/ddddemo/api/http/middleware"
	"github.com/jettjia/ddddemo/api/http/router"
	"github.com/jettjia/ddddemo/config"
)

func InitHttp() {
	// open api
	go func() {
		runPublicHttp()
	}()

}

// 外部接口
func runPublicHttp() {
	var engine *gin.Engine

	cfg := config.NewConfig()
	engineDefault := gin.Default()
	gin.SetMode(cfg.Server.Mode)

	// 配置跨域
	engineDefault.Use(middleware.Cors())
	// 全局recover
	engineDefault.Use(middleware.CatchError())
	// 通用中间件
	engineDefault.Use(middleware.Universal())
	// otel
	if cfg.Otel.Enable {
		// set otel config
		os.Setenv("OTEL_OTLP_HTTP_ENDPOINT", cfg.Otel.ExportEndpoint)
		os.Setenv("OTEL_USERNAME", cfg.Otel.Username)
		os.Setenv("OTEL_PASSWORD", cfg.Otel.Password)
		os.Setenv("OTEL_SERVICE_NAME", cfg.Server.ServerName)

		tp := tel.InitTracerHTTP()
		defer func() {
			if err := tp.Shutdown(context.Background()); err != nil {
				fmt.Println("Error shutting down tracer provider: ", err)
			}
		}()
		engineDefault.Use(otelgin.Middleware(cfg.Server.ServerName))
	}

	// 健康检查
	engineDefault.GET("/health/ready", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "ready")
	})
	engineDefault.GET("/health/alive", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.String(http.StatusOK, "alive")
	})

	engine = router.Routers(engineDefault)
	err := engine.Run(fmt.Sprintf(":%d", cfg.Server.PublicPort)) // 启动web
	if err != nil {
		panic(err)
	}
}
