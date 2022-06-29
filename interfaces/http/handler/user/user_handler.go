package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dto "github.com/jettjia/go-ddd/application/dto/user"
	service "github.com/jettjia/go-ddd/application/service/user"
	"github.com/jettjia/go-ddd/infrastructure/pkg/response"
)

type UserHandler struct {
	UserSrv *service.UserService
}

// ApiGetSimpleUser 查询简单的用户信息 GET /user/123
func (u *UserHandler) ApiGetSimpleUser(c *gin.Context) {
	simpleUserReq := &dto.SimpleUserInfoReq{}
	c.BindJSON(simpleUserReq) // 处理请求参数

	dtoSimpleUserInfo, err := u.UserSrv.GetSimpleUserInfo(c.Request.Context(), simpleUserReq)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(dtoSimpleUserInfo))

	return
}
