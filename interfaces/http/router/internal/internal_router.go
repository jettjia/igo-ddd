package router

import (
	"github.com/gin-gonic/gin"

	"github.com/jettjia/go-ddd-demo/cmd"
	"github.com/jettjia/go-ddd-demo/interfaces/http/handler/internal_handler"
)

func InitInternalRouter(Router *gin.RouterGroup, app *cmd.App) {
	hand := handler.InternalHandler{}
	router := Router.Group("/sys")
	{
		// demo
		router.GET("/demo", hand.InternalDemoFunc)
	}
}
