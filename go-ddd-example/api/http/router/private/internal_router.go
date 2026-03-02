package private

import (
	"github.com/gin-gonic/gin"

	"github.com/jettjia/ddddemo/api/http/handler/private"
)

func SetPrivateRouter(Router *gin.RouterGroup) {
	handUser := private.NewPrivateHandler()

	GRouter := Router.Group("/example")
	{
		GRouter.POST("/func1", handUser.Internal) // 内部接口
	}
}
