package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"

	"jettjia/go-ddd-demo-multi-common/pkg/response"
	"jettjia/go-ddd-demo-multi-system/config"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

func TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		conf := config.NewConfig()
		if conf.Server.Dev {
			return
		}

		auth := c.GetHeader("Authorization")
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == "" {
			err := gerror.NewCode(response.CommUnauthorized, "Unauthorized")
			response.RspErr(c, err)
			c.Abort()
			return
		}

		// 校验token
		claims, err := entity.ParseToken(token)
		if err != nil {
			err = gerror.NewCode(response.CommUnauthorized, "Unauthorized")
			response.RspErr(c, err)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			err = gerror.NewCode(response.CommUnauthorized, "Token expired")
			response.RspErr(c, err)
			c.Abort()
			return
		}
	}
}
