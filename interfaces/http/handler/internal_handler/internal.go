package internal_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"net/http"
)

type InternalHandler struct {
}

func (h *InternalHandler) InternalDemoFunc(c *gin.Context) {
	responseutil.RspOk(c, http.StatusOK, "internal demo func")
}
