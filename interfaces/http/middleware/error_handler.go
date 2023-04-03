package middleware

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	for _, err := range c.Errors {
		// 统一处理 mysql 1062 错误，sql内容冲突
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			responseutil.RspErr(c, gerror.NewCode(responseutil.CommConflict, err.Error()))
			return
		}

		// 统一处理 mysql 1054 错误,sql字段错误
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1054 {
			responseutil.RspErr(c, gerror.NewCode(responseutil.CommForbidden, err.Error()))
			return
		}
		// 统一处理 gorm NotFound
		if err.Error() == "record not found" {
			responseutil.RspErr(c, gerror.NewCode(responseutil.CommNotFound, err.Error()))
			return
		}

		global.GLog.Errorln(fmt.Sprintf("%+v", err)) // 记录到日志
		responseutil.RspErr(c, err)                  // 接口错误
	}
}
