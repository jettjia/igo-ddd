package user

import (
	"github.com/gin-gonic/gin"

	handler "github.com/jett/gin-ddd/interfaces/http/handler/user"
	"github.com/jett/gin-ddd/interfaces/http/registersrv"
)

func InitUserRouter(Router *gin.RouterGroup, registerSrv *registersrv.Registersrv) {
	hand := handler.UserHandler{
		UserSrv: registerSrv.UserSrv,
	}
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("getSimpleUserInfo", hand.ApiGetSimpleUser) // 查找用户
	}
}
