package public

import (
	"github.com/gin-gonic/gin"

	handler "jettjia/go-ddd-demo-multi-system/interface/http/handler/public"
)

func SetPublicRouter(Router *gin.RouterGroup, hand *handler.Handler) {
	UserRouter := Router.Group("/sys")
	{
		UserRouter.POST("/menu", hand.CreateSysMenu)         // 创建
		UserRouter.DELETE("/menu/:ulid", hand.DeleteSysMenu) // 删除
		UserRouter.PUT("/menu/:ulid", hand.UpdateSysMenu)    // 修改
		UserRouter.GET("/menu/:ulid", hand.FindSysMenuById)  // 查询ByID
		UserRouter.POST("/menuPage", hand.FindSysMenuPage)   // 查询分页
	}
}
