package internal

import (
	"github.com/gin-gonic/gin"
	handler "github.com/jettjia/go-ddd-demo/interfaces/http/handler/internal_handler"
)

func InitInternalRouter(Router *gin.RouterGroup) {
	hand := handler.InternalHandler{}
	router := Router.Group("/sys")
	{
		// demo
		router.GET("/demo", hand.InternalDemoFunc)
	}
}
