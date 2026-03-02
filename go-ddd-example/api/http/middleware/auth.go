package middleware

import (
	"github.com/gin-gonic/gin"

	middleware "github.com/jettjia/igo-pkg/pkg/xmiddleware/mgin"

	"github.com/jettjia/ddddemo/config"
)

func TokenAuthorization() gin.HandlerFunc {
	conf := config.NewConfig()

	return middleware.TokenAuthorizationHydra(conf)
}
