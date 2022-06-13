package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

}
