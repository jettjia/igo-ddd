package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Otelgin() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			otelgin.Middleware("go-ddd-demo")
		}()
		c.Next()
	}
}
