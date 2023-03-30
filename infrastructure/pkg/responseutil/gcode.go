package responseutil

import "github.com/gogf/gf/v2/errors/gcode"

type BizCode struct {
	code    int
	message string
	detail  BizCodeDetail
}

type BizCodeDetail struct {
	SubCode int
}

func (c BizCode) BizDetail() BizCodeDetail {
	return c.detail
}

func (c BizCode) Code() int {
	return c.code
}

func (c BizCode) Message() string {
	return c.message
}

func (c BizCode) Detail() interface{} {
	return c.detail
}

func NewCode(httpCode int, subCode int, message string) gcode.Code {
	return BizCode{
		code:    httpCode,
		message: message,
		detail: BizCodeDetail{
			SubCode: subCode,
		},
	}
}
