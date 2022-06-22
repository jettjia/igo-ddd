package user

import (
	"github.com/gin-gonic/gin"

	"github.com/jett/gin-ddd/cmd"
	handler "github.com/jett/gin-ddd/interfaces/http/handler/user"
)

func InitUserRouter(Router *gin.RouterGroup, app *cmd.App) {
	hand := handler.UserHandler{
		UserSrv: app.UserSrv,
	}
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("getSimpleUserInfo", hand.ApiGetSimpleUser) // 查找用户
	}
}
