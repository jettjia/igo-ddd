package middleware

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/jettjia/igo-pkg/pkg/xmiddleware/mgin"
)

// Universal 通用的中间件部分
// 比如处理用户登录的信息
func Universal() gin.HandlerFunc {
	return middleware.Universal()
}
