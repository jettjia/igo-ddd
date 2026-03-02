package gen

const TempInterfacesRouter = `
package public

import (
	"github.com/gin-gonic/gin"

	"xdata/xtext/zsvc-example/api/http/handler/public"
	handUser "xdata/xtext/zsvc-example/api/http/handler/public/user"
)

func SetPublicRouter(Router *gin.RouterGroup) {
	handUser := handUser.NewHandler()

	GRouter := Router.Group("/user")
	{
		// menu
		GRouter.POST("/menu", handUser.CreateSysMenu)              // 创建
		GRouter.DELETE("/menu/:ulid", handUser.DeleteSysMenu)      // 删除
		GRouter.PUT("/menu/:ulid", handUser.UpdateSysMenu)         // 修改
		GRouter.GET("/menu/:ulid", handUser.FindSysMenuById)       // 查询ByID
		GRouter.POST("/menuPage", handUser.FindSysMenuPage)        // 查询分页
	}
}
`
