package log

import (
	"github.com/sirupsen/logrus"

	"github.com/jettjia/igo-pkg/pkg/log"

	"github.com/jettjia/ddddemo/config"
)

func NewLogger() *logrus.Logger {
	conf := config.NewConfig()

	return log.NewLoggerServer(conf)
}
