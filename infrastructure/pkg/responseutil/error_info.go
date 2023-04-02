package responseutil

import (
	"path/filepath"
	"runtime"
	"strings"
)

// ErrorInfo 错误信息
type ErrorInfo struct {
	Internal []frameErrorInfo
}

type frameErrorInfo struct {
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	FuncName string `json:"func_name"`
}

// Panic 异常
func Panic() ErrorInfo {
	return alarm()
}

// GrpcPanic 异常
func GrpcPanic() ErrorInfo {
	return alarm()
}

func SqlError() ErrorInfo {
	return alarm()
}

func alarm() (err ErrorInfo) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	n := runtime.Callers(0, pc)
	frames := runtime.CallersFrames(pc[:n])

	var frameErrorInfosInternal []frameErrorInfo

	for {
		frame, more := frames.Next()
		if strings.Contains(frame.File, "runtime/") {
			continue
		}
		if strings.Contains(frame.File, "gin-gonic/") {
			continue
		}

		// 记录具体的错误到日日志中
		var (
			frameErrorInfo frameErrorInfo
		)

		_, fileName := filepath.Split(frame.File)
		frameErrorInfo.Filename = fileName
		frameErrorInfo.Line = frame.Line
		frameErrorInfo.FuncName = frame.Function
		frameErrorInfosInternal = append(frameErrorInfosInternal, frameErrorInfo)

		if !more {
			break
		}
	}

	err.Internal = frameErrorInfosInternal
	return
}
