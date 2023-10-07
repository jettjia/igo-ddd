package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"jettjia/go-ddd-demo-multi-common/pkg/log"
	"jettjia/go-ddd-demo-multi-common/pkg/response"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 捕获panic错误,比如 panic()
			if errAny := recover(); errAny != nil {
				// 记录到log
				log.NewLogger().WithFields(logrus.Fields{"url": c.Request.URL.Path}).
					WithFields(logrus.Fields{"method": c.Request.Method}).
					Errorln(errAny)

				switch errAny.(type) {
				case error:
					// 统一处理 mysql 1062 错误，sql内容冲突
					var mysqlErr *mysql.MySQLError
					if errors.As(errAny.(error), &mysqlErr) && mysqlErr.Number == 1062 {
						response.RspErr(c, gerror.NewCode(response.CommConflict, mysqlErr.Message))
						c.Abort()
						return
					}

					// 统一处理 mysql 1054 错误,sql字段错误
					if errors.As(errAny.(error), &mysqlErr) && mysqlErr.Number == 1054 {
						response.RspErr(c, gerror.NewCode(response.CommForbidden, mysqlErr.Message))
						c.Abort()
						return
					}

					// string错误
					if errAny != "" {
						errorInfo := response.Panic(errAny)
						response.RspErr(c, gerror.NewCode(response.CommInternalServer, fmt.Sprintf("errMsg:%+v; errStack:%+v", errAny, errorInfo.Internal))) // 前端返回
						c.Abort()
						return
					}

				default:
					// 统一处理 其他错误
					errorInfo := response.Panic(errAny)
					response.RspErr(c, gerror.NewCode(response.CommInternalServer, fmt.Sprintf("errMsg:%+v; errStack:%+v", errAny, errorInfo.Internal))) // 前端返回
					c.Abort()
					return
				}
			}

			// 手动抛出的错误,比如 errors.New()
			if len(c.Errors) != 0 {
				for _, errAny := range c.Errors {
					// 统一处理 gorm NotFound
					if errAny.Error() == "Record does not exist" {
						response.RspErr(c, gerror.NewCode(response.CommNotFound, errAny.Error()))
						return
					}

					// 统一处理 其他错误
					errorInfo := response.Panic(errAny)
					response.RspErr(c, gerror.NewCode(response.CommInternalServer, fmt.Sprintf("errMsg:%+v; errStack:%+v", errAny, errorInfo.Internal))) // 前端返回
					c.Abort()
					return
				}
			}
		}()
		c.Next()
	}
}
