package responseutil

import "net/http"

// 通用异常
var (
	BadRequest          = 400000000
	Unauthorized        = 401000000
	Forbidden           = 403000000
	NotFound            = 404000000
	Conflict            = 409000000
	InternalServerError = 500000000
)

var (
	// CommErr sap链接配置模块
	CommErr               = 100000
	CommErrBadRequest     = BadRequest + CommErr
	CommErrUnauthorized   = Unauthorized + CommErr
	CommErrForbidden      = Forbidden + CommErr
	CommErrNotFound       = NotFound + CommErr
	CommErrConflict       = Conflict + CommErr
	CommErrInternalServer = InternalServerError + CommErr
)

var (
	CommBadRequest     = NewCode(http.StatusBadRequest, CommErrBadRequest, "Invalid parameter")
	CommUnauthorized   = NewCode(http.StatusUnauthorized, CommErrUnauthorized, "Unauthorized")
	CommForbidden      = NewCode(http.StatusForbidden, CommErrForbidden, "Forbidden")
	CommNotFound       = NewCode(http.StatusNotFound, CommErrNotFound, "Record does not exist")
	CommConflict       = NewCode(http.StatusConflict, CommErrConflict, "The request is conflicts")
	CommInternalServer = NewCode(http.StatusInternalServerError, CommErrInternalServer, "Internal Server Error")
)
