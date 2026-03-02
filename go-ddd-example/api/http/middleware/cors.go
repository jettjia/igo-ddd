package middleware

import (
	middleware "github.com/jettjia/igo-pkg/pkg/xmiddleware/mgin"

	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return middleware.Cors()
}
