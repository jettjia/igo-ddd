package user

import (
	"github.com/gin-gonic/gin"
	handler "github.com/jett/gin-ddd/interfaces/http/handler/user"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userHand := handler.UserHandler{} // todo 这里有问题，service没有注入
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("getSimpleUserInfo", userHand.ApiGetSimpleUser) // 查找用户
	}
}
