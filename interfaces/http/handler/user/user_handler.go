package handler

import (
	"github.com/gin-gonic/gin"

	dto "github.com/jett/gin-ddd/application/dto/user"
	service "github.com/jett/gin-ddd/application/service/user"
	"github.com/jett/gin-ddd/interfaces/http/handler"
)

type UserHandler struct {
	userSrv *service.UserService
}

// ApiGetSimpleUser 查询简单的用户信息 GET /user/123
func (this *UserHandler) ApiGetSimpleUser(ctx *gin.Context) {
	simpleUserReq := &dto.SimpleUserInfoReq{}
	ctx.BindJSON(simpleUserReq) // 处理请求参数

	dtoSimpleUserInfo := this.userSrv.GetSimpleUserInfo(simpleUserReq)

	handler.ResponseData(ctx, dtoSimpleUserInfo)

	return
}

// ApiSaveUser 保存用户
func ApiSaveUser(ctx *gin.Context) {
}
