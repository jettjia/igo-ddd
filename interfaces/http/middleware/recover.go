package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorInfo := responseutil.Panic(fmt.Sprintf("%s", err))
				global.GLog.Errorln(util.PrintJson(errorInfo.Internale))                                             // 记录到日志
				responseutil.RspErr(c, gerror.NewCode(responseutil.CommInternalServer, gconv.String(errorInfo.Out))) // 前端返回
				return
			}
		}()
		c.Next()
	}
}
