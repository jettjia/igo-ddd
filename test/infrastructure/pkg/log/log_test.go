package log

import (
	"testing"

	log2 "github.com/jettjia/gin-ddd/infrastructure/pkg/log"
	_ "github.com/jettjia/gin-ddd/test"
)

func TestLog(t *testing.T) {
	log := log2.NewLogger()
	log.Errorf("ddddd")
}
