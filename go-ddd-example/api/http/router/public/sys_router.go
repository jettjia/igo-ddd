package public

import (
	"github.com/gin-gonic/gin"

	handUser "github.com/jettjia/ddddemo/api/http/handler/public/user"
)

func SetPublicRouter(Router *gin.RouterGroup) {
	handUser := handUser.NewHandler()

	GRouter := Router.Group("/user")
	{

		// user
		GRouter.POST("/user", handUser.CreateSysUser)              // 创建
		GRouter.DELETE("/user/:ulid", handUser.DeleteSysUser)      // 删除
		GRouter.PUT("/user/:ulid", handUser.UpdateSysUser)         // 修改
		GRouter.GET("/user/:ulid", handUser.FindSysUserById)       // 查询ByID
		GRouter.POST("/user/byQuery", handUser.FindSysUserByQuery) // 查询ByQuery
		GRouter.POST("/user/byAll", handUser.FindSysUserAll)       // 查询ByAll
		GRouter.POST("/userPage", handUser.FindSysUserPage)        // 查询分页
	}
}
