package log

import (
	"io"
	"os"
	"path"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerConfig struct {
	LogPath    string // 日志路径
	MaxSize    int    // 日志文件大小，单位是 MB
	MaxBackups int    // 最大过期日志保留个数
	MaxAge     int    // 保留过期文件最大时间，单位 天
	LogLevel   string // 日志级别 panic, fatal, error, warn, info, debug,trace
	LogOut     string // console, file
}

var (
	once sync.Once
	l    *logrus.Logger
)

func NewLogger(options ...func(*LoggerConfig)) *logrus.Logger {
	once.Do(func() {
		cfg := LoggerConfig{
			LogPath:    "/tmp/logs/",
			MaxSize:    128,
			MaxBackups: 255,
			MaxAge:     7,
			LogLevel:   "error",
			LogOut:     "console",
		}

		for _, option := range options {
			option(&cfg)
		}

		l = initLogger(&cfg)
	})
	return l
}

func initLogger(cfg *LoggerConfig) *logrus.Logger {
	logHandle := logrus.New()
	logHandle.SetLevel(logLevel(cfg.LogLevel))
	logHandle.SetFormatter(&logrus.JSONFormatter{})

	var output io.Writer
	if cfg.LogOut == "console" {
		output = os.Stdout
	} else {
		output = &lumberjack.Logger{
			Filename:   logFileNamePath(cfg.LogPath),
			MaxSize:    cfg.MaxSize,    // 日志文件大小，单位是 MB
			MaxBackups: cfg.MaxBackups, // 最大过期日志保留个数
			MaxAge:     cfg.MaxAge,     // 保留过期文件最大时间，单位 天
			Compress:   true,           // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
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

func logFileNamePath(settingPath string) string {
	var (
		logFilePath string
	)
	logFilePath = settingPath
	if logFilePath == "" {
		logFilePath = "/tmp/logs/"
	}

	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	return fileName
}
