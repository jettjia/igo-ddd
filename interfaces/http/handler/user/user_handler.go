package user

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	dto "github.com/jett/gin-ddd/application/dto/user"
	service "github.com/jett/gin-ddd/application/service/user"
	"github.com/jett/gin-ddd/infrastructure/pkg/response"
)

// RESTHandler 公共RESTful api Handler接口
type RESTHandler interface {
	// 注册开放API
	RegisterAPI(engine *gin.Engine)
}

var (
	once sync.Once
	h    RESTHandler
)

// NewRESTHandler 创建公共RESTful api handler对象
func NewRESTHandler() RESTHandler {
	once.Do(func() {
		h = &restHandler{
			userSrv: service.NewUserService(),
		}
	})

	return h
}

type restHandler struct {
	userSrv *service.UserService
}

// RegisterAPI 注册开放API
func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/user/findOne", h.ApiGetSimpleUser)
}

// ApiGetSimpleUser 查询简单的用户信息
func (h *restHandler) ApiGetSimpleUser(c *gin.Context) {
	simpleUserReq := &dto.SimpleUserInfoReq{}
	c.BindJSON(simpleUserReq) // 处理请求参数
	dtoSimpleUserInfo, err := h.userSrv.GetSimpleUserInfo(simpleUserReq)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(dtoSimpleUserInfo))

	return
}
