package log

import (
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/jett/gin-ddd/infrastructure/config"
)

// Logger 服务日志服务，可适配其他日志组件
type Logger interface {
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})
	Debugf(format string, args ...interface{})
	Debugln(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
	Tracef(format string, args ...interface{})
	Traceln(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
}

var (
	logOnce sync.Once
	l       Logger
)

// NewLogger 获取日志句柄
func NewLogger() Logger {
	logOnce.Do(initLogger)

	return l
}

type serverLog struct {
	logger *logrus.Logger
}

// Infof 普通信息
func (l *serverLog) Infof(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Infof(format, args...)
}

// Infoln 普通信息
func (l *serverLog) Infoln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Infoln(args...)
}

// Warnf 警告信息
func (l *serverLog) Warnf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Warnf(format, args...)
}

// Warnln 警告信息
func (l *serverLog) Warnln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Warnln(args...)
}

// Errorf 错误信息
func (l *serverLog) Errorf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Errorf(format, args...)
}

// Errorln 错误信息
func (l *serverLog) Errorln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Errorln(args...)
}

// Debugf 调试信息
func (l *serverLog) Debugf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debugf(format, args...)
}

// Debugln 调试信息
func (l *serverLog) Debugln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Debugln(args...)
}

// Tracef 跟踪信息
func (l *serverLog) Tracef(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Tracef(format, args...)
}

// Traceln 跟踪信息
func (l *serverLog) Traceln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Traceln(args...)
}

// Fatalf 致命错误
func (l *serverLog) Fatalf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Fatalf(format, args...)
}

// Fatalln 致命错误
func (l *serverLog) Fatalln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Fatalln(args...)
}

// Panicf 恐慌错误
func (l *serverLog) Panicf(format string, args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Panicf(format, args...)
}

// Panicln 恐慌错误
func (l *serverLog) Panicln(args ...interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.Panicln(args...)
}

func initLogger() {
	logHandle := &serverLog{}
	logHandle.logger = logrus.New()
	logHandle.logger.SetFormatter(&logrus.JSONFormatter{})

	conf := config.NewConfig() // 读取配置

	logDir := conf.Log.LogFileDir
	if logDir == "" {
		logDir = "/tmp/log/"
	}
	logName := conf.Log.AppName
	if logName == "" {
		logName = "server.log"
	}

	logFileName := logDir + "/" + logName + ".log"

	loggerOut := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    conf.Log.MaxSize,    // 日志文件大小，单位是 MB
		MaxBackups: conf.Log.MaxBackups, // 最大过期日志保留个数
		MaxAge:     conf.Log.MaxAge,     // 保留过期文件最大时间，单位 天
		Compress:   true,                // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}

	logHandle.logger.SetOutput(loggerOut)

	l = logHandle
}
