package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/jwt"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
)

func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		return // 不校验token
		tokenID := c.GetHeader("Authorization")
		token := strings.TrimPrefix(tokenID, "Bearer ")
		if token == "" {
			err := gerror.NewCode(responseutil.CommUnauthorized, "Unauthorized")
			responseutil.RspErr(c, err)
			c.Abort()
			return
		}

		// 校验token
		claims, err := entity.ParseToken(token)
		if err != nil {
			err := gerror.NewCode(responseutil.CommUnauthorized, "Unauthorized")
			responseutil.RspErr(c, err)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			err := gerror.NewCode(responseutil.CommUnauthorized, "Token expired")
			responseutil.RspErr(c, err)
			c.Abort()
			return
		}

		global.GCustomerInfo = claims.CustomerInfo
	}
}
