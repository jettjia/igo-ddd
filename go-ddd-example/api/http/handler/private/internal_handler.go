package private

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *PrivateHandler) Internal(c *gin.Context) {
	fmt.Println("internal")
}
