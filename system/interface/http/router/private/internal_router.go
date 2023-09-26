package private

import (
	"github.com/gin-gonic/gin"

	handler "jettjia/go-ddd-demo-multi-system/interface/http/handler/private"
)

func SetPrivateRouter(Router *gin.RouterGroup, hand *handler.PrivateHandler) {
	UserRouter := Router.Group("/sys")
	{
		UserRouter.GET("/demo", hand.Demo) // demo
	}
}
