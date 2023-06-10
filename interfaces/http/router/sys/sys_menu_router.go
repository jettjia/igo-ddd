package sys

import (
	"github.com/gin-gonic/gin"
	handler "github.com/jettjia/go-ddd-demo/interfaces/http/handler/sys"
)

func InitSysRouter(Router *gin.RouterGroup) {
	hand := handler.NewSysMenuHandler()
	UserRouter := Router.Group("/sys")
	{
		UserRouter.POST("menu", hand.CreateSysMenu)       // 创建
		UserRouter.DELETE("menu/:id", hand.DeleteSysMenu) // 删除
		UserRouter.PUT("menu/:id", hand.UpdateSysMenu)    // 修改
		UserRouter.GET("menu/:id", hand.FindSysMenuById)  // 查询ByID
		UserRouter.POST("menuPage", hand.FindSysMenuPage) // 查询分页
	}
}
