package health

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
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
		h = &restHandler{}
	})

	return h
}

type restHandler struct {
}

// RegisterAPI 注册开放API
func (h *restHandler) RegisterAPI(engine *gin.Engine) {
	engine.GET("/health/ready", h.getHealth)
	engine.GET("/health/alive", h.getAlive)
}

func (h *restHandler) getHealth(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "ready")
}

func (h *restHandler) getAlive(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, "alive")
}
