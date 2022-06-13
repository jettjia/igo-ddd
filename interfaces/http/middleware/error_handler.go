package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

}
