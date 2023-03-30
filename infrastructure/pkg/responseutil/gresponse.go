package responseutil

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/types"
)

type rspCreateData struct {
	Id interface{} `json:"id"`
}

type rspErrorData struct {
	Cause   interface{} `json:"cause"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

type rspListNull struct {
	Entries  []int          `json:"entries"`
	PageData types.PageData `json:"page_data"`
}

// RspOk 返回操作成功
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
			rspCreateData{
				Id: any,
			},
		)
	case 204:
		c.AbortWithStatus(
			http.StatusNoContent,
		)
	case 200200:
		rspNull := rspListNull{
			Entries: make([]int, 0),
		}
		c.JSON(
			http.StatusOK,
			rspNull,
		)
	case 200201:
		c.JSON(
			http.StatusOK,
			make([]int, 0),
		)

	default:
		c.JSON(
			http.StatusOK,
			any,
		)
	}
}

// RspErr 返回操作失败
func RspErr(c *gin.Context, err error) {
	code := gerror.Code(err)
	if code == gcode.CodeNil && err != nil {
		code = gcode.CodeInternalError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = gerror.NewCode(CommNotFound, err.Error())
		} else {
			err = gerror.NewCode(CommInternalServer, err.Error())
		}
		code = gerror.Code(err)
	}
	subCode := code.(BizCode).BizDetail().SubCode
	msg := msg(code.Code())
	rspError(c, code.Code(), subCode, msg, err)
}

func rspError(c *gin.Context, code int, subCode int, message string, err error) {
	c.JSON(
		code,
		rspErrorData{
			Cause:   err.Error(),
			Code:    subCode,
			Message: message,
		},
	)
}

func msg(code int) (msg string) {
	lang := global.Gconfig.Server.Lang
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
