// Package responseutil 捕获错误的堆栈信息
package responseutil

import (
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/jettjia/go-ddd-demo/global"
)

// ErrorInfo 堆栈错误信息
type (
	ErrorInfo struct {
		Internal []frameErrorInfo `json:"stacks"`
	}

	frameErrorInfo struct {
		Filename string `json:"filename"`
		Line     int    `json:"line"`
		FuncName string `json:"func_name"`
	}
)

// 记录到日志的错误信息
type logErr struct {
	Stack []frameErrorInfo `json:"stack"`
	Sql   string           `json:"sql,omitempty"`
}

// Panic 异常
func Panic(err interface{}) ErrorInfo {
	var logErr logErr
	errInfo := alarm()
	logErr.Stack = errInfo.Internal
	global.GLog.WithFields(logrus.Fields{
		"detail": logErr,
	}).Errorln(err) // 记录到日志

	return errInfo
}

// GrpcPanic 异常
func GrpcPanic(err string) ErrorInfo {
	var logErr logErr
	errInfo := alarm()
	logErr.Stack = errInfo.Internal
	global.GLog.WithFields(logrus.Fields{
		"detail": logErr,
	}).Errorln(err) // 记录到日志

	return errInfo
}

// SqlError 异常
func SqlError(err string, sql string) {
	var logErr logErr
	errInfo := alarm()
	logErr.Stack = errInfo.Internal
	logErr.Sql = sql
	global.GLog.WithFields(logrus.Fields{
		"detail": logErr,
	}).Errorln(err) // 记录到日志
	return
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
		if strings.Contains(frame.File, "infrastructure/pkg") {
			continue
		}

		// 记录具体的错误到日日志中
		var (
			frameErrorInfo frameErrorInfo
		)

		_, fileName := filepath.Split(frame.File)
		frameErrorInfo.Filename = fileName
		frameErrorInfo.Line = frame.Line
		_, function := filepath.Split(frame.Function)
		frameErrorInfo.FuncName = function
		frameErrorInfosInternal = append(frameErrorInfosInternal, frameErrorInfo)

		if !more {
			break
		}
	}

	err.Internal = frameErrorInfosInternal
	return
}
