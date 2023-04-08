package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorInfo := responseutil.Panic(err)
				responseutil.RspErr(c, gerror.NewCodef(responseutil.CommInternalServer, "%+v", errorInfo.Internal)) // 前端返回
				return
			}
		}()
		c.Next()
	}
}
