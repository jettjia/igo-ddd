package config

import (
	"fmt"
	"testing"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/global"
)

func TestNewConfig(t *testing.T) {

	if global.Gconfig.Server.Address == "" {
		t.Fatal("error: 请在 config.yaml 中配置端口号")
	}
	fmt.Println(global.Gconfig.Server.Address)
}
