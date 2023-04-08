package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/jettjia/go-ddd-demo/infrastructure/config"
)

// NewLogger 获取日志句柄
func NewLogger(env string) *logrus.Logger {
	return initLogger(env)
}

func initLogger(env string) *logrus.Logger {
	conf := config.NewConfig(env) // 读取配置

	logHandle := logrus.New()
	logHandle.SetLevel(logLevel(conf.Log.LogLevel))
	logHandle.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logDir := conf.Log.LogFileDir
	if logDir == "" {
		logDir = "/tmp/log/"
	}
	logName := conf.Log.AppName
	if logName == "" {
		logName = "server.log"
	}

	logFileName := logDir + "/" + logName + ".log"

	var output io.Writer
	if conf.Log.LogOut == "console" {
		output = os.Stdout
	} else {
		output = &lumberjack.Logger{
			Filename:   logFileName,
			MaxSize:    conf.Log.MaxSize,    // 日志文件大小，单位是 MB
			MaxBackups: conf.Log.MaxBackups, // 最大过期日志保留个数
			MaxAge:     conf.Log.MaxAge,     // 保留过期文件最大时间，单位 天
			Compress:   true,                // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
		}
	}

	logHandle.SetOutput(output)

	return logHandle
}

func logLevel(logLevel string) (level logrus.Level) {
	switch logLevel {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	}

	return logrus.DebugLevel
}
