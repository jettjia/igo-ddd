package private

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jettjia/go-ddd-demo-multi-common/pkg/log"
)

func (h *PrivateHandler) Demo(c *gin.Context) {
	log.NewLogger().Infoln("aitext-go-ddd") // 自定义Log,建议使用
	c.JSON(http.StatusOK, "i am internal api")
}
