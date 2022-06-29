package config

import (
	"fmt"
	"testing"

	"github.com/jettjia/gin-ddd/global"
	_ "github.com/jettjia/gin-ddd/test"
)

func TestNewConfig(t *testing.T) {

	if global.Gconfig.Server.Address == "" {
		t.Fatal("error: 请在 config.yaml 中配置端口号")
	}
	fmt.Println(global.Gconfig.Server.Address)
}
