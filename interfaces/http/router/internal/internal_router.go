package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/jettjia/go-ddd-demo/boot"

	handler "github.com/jettjia/go-ddd-demo/interfaces/http/handler/internal_handler"
)

func InitInternalRouter(Router *gin.RouterGroup, app *boot.App) {
	hand := handler.InternalHandler{}
	router := Router.Group("/sys")
	{
		// demo
		router.GET("/demo", hand.InternalDemoFunc)
	}
}
