package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/jettjia/ddddemo/types/apierror"
	"github.com/jettjia/igo-pkg/pkg/xerror"
	"github.com/jettjia/igo-pkg/pkg/xresponse"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			// 捕获panic错误,比如 panic()
			if errAny := recover(); errAny != nil {
				switch errAny.(type) {
				case error:
					// 统一处理 mysql 1062 错误，sql内容冲突
					var mysqlErr *mysql.MySQLError
					if errors.As(errAny.(error), &mysqlErr) && mysqlErr.Number == 1062 {
						err := xerror.NewErrorOpt(apierror.ConflictErr, xerror.WithCause(mysqlErr.Message))
						xresponse.RspErr(c, err)
						c.Abort()
						return
					}

					// 统一处理 mysql 1054 错误,sql字段错误
					if errors.As(errAny.(error), &mysqlErr) && mysqlErr.Number == 1054 {
						err := xerror.NewErrorOpt(apierror.ForbiddenErr, xerror.WithCause(mysqlErr.Message))
						xresponse.RspErr(c, err)
						c.Abort()
						return
					}

					// 统一处理 mysql 1064 错误,sql语法错误，比如多个引号，少个括号
					if errors.As(errAny.(error), &mysqlErr) && mysqlErr.Number == 1064 {
						err := xerror.NewErrorOpt(apierror.ForbiddenErr, xerror.WithCause(mysqlErr.Message))
						xresponse.RspErr(c, err)
						c.Abort()
						return
					}

				default:
					// 统一处理 其他错误
					err := xerror.NewErrorOpt(apierror.InternalServerErr, xerror.WithCause(fmt.Sprintf("%+v", errAny)))
					xresponse.RspErr(c, err)
					c.Abort()
					return
				}
			}

			// 手动抛出的错误,比如 ierror.New()
			if len(c.Errors) != 0 {
				// xerror.RecoverWithStack() // print error
				for _, errAny := range c.Errors {
					switch errAny.Error() {
					case "EOF":
						err := xerror.NewErrorOpt(apierror.BadRequestErr, xerror.WithCause(errAny.Error()))
						xresponse.RspErr(c, err)
						c.Abort()
						return
					default:
						// 统一处理 其他错误
						xresponse.RspErr(c, errAny)
						c.Abort()
						return
					}
				}
			}
		}()
		c.Next()
	}
}
