package log

import (
	"testing"

	log2 "github.com/jett/gin-ddd/infrastructure/pkg/log"
)

func TestLog(t *testing.T) {
	log := log2.NewLogger()
	log.Errorf("ddddd")
}
