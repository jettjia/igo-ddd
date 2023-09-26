package response

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
)

type rspErrorData struct {
	SubCode int         `json:"sub_code"`
	Message string      `json:"message"`
	Cause   interface{} `json:"cause"`
}

func RspOk(c *gin.Context, code int, any interface{}) {
	switch code {
	case 200:
		c.JSON(
			http.StatusOK,
			any,
		)
	case 201:
		c.JSON(
			http.StatusCreated,
			any,
		)
	case 204:
		c.AbortWithStatus(
			http.StatusNoContent,
		)

	default:
		c.JSON(
			http.StatusOK,
			any,
		)
	}
}

func RspErr(c *gin.Context, err error) {
	code := gerror.Code(err)
	subCode := code.(BizCode).BizDetail().SubCode
	msg := msg(code.Code())
	rspError(c, code.Code(), subCode, msg, err)
}

func rspError(c *gin.Context, code int, subCode int, message string, err error) {
	c.JSON(
		code,
		rspErrorData{
			SubCode: subCode, // error sub code
			Message: message, // error message
			Cause:   err,     // stack error
		},
	)
}

func msg(code int) (msg string) {
	lang := "en"
	i18n := gi18n.New()
	ctx := gi18n.WithLanguage(context.TODO(), lang)

	switch code {
	case http.StatusBadRequest:
		msg = i18n.Translate(ctx, "{#BadRequest}")
	case http.StatusUnauthorized:
		msg = i18n.Translate(ctx, "{#Unauthorized}")
	case http.StatusForbidden:
		msg = i18n.Translate(ctx, "{#Forbidden}")
	case http.StatusNotFound:
		msg = i18n.Translate(ctx, "{#NotFound}")
	case http.StatusConflict:
		msg = i18n.Translate(ctx, "{#Conflict}")
	case http.StatusInternalServerError:
		msg = i18n.Translate(ctx, "{#InternalServerError}")
	}

	return
}
