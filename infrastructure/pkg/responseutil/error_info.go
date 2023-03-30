package responseutil

import (
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
)

// ErrorInfo 错误信息
type ErrorInfo struct {
	Internale ErrorInfoInternal `json:"internale"`
	Out       ErrorInfoOut      `json:"out"`
}

type ErrorInfoInternal struct {
	Message  string `json:"message"`
	Level    string `json:"level"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

type ErrorInfoOut struct {
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

// Panic 异常
func Panic(errText string) ErrorInfo {
	debug.PrintStack()
	return alarm("PANIC", errText, 5)
}

// GrpcPanic 异常
func GrpcPanic(errText string) ErrorInfo {
	return alarm("GrpcPanic", errText, 7)
}

func SqlError(errText string) ErrorInfo {
	return alarm("ERROR", errText, 5)
}

func alarm(level string, errText string, skip int) ErrorInfo {
	// 定义 文件名、行号、方法名
	fileName, line, functionName, fileNameShort := "?", 0, "?", "?"
	pc, fileName, line, ok := runtime.Caller(skip)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")

		_, fileNameShort = filepath.Split(fileName)
	}

	var msg = ErrorInfo{
		Internale: ErrorInfoInternal{
			Level:    level,
			Message:  errText,
			Filename: fileName,
			Line:     line,
			Funcname: functionName,
		},

		Out: ErrorInfoOut{
			Message:  errText,
			Filename: fileNameShort,
			Line:     line,
			Funcname: functionName,
		},
	}

	return msg
}
