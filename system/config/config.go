package config

import (
	"flag"
	"os"
	"sync"

	"github.com/google/wire"

	"jettjia/go-ddd-demo-multi-common/pkg/conf"
)

var CfgProvider = wire.NewSet(NewConfig)

func NewConfig() (conf *conf.Config) {
	conf = initConfig()
	return
}

var (
	configOnce sync.Once
	cfg        *conf.Config
)

// InitConfig 读取配置
func initConfig() *conf.Config {
	env := os.Getenv("env")

	if len(env) == 0 {
		env = "debug"
	}
	configOnce.Do(func() {

		file := "./manifest/config/config-" + env + ".yaml"

		cfg = &conf.Config{}
		flag.Set("conf", file)
		if err := conf.ParseYaml(cfg); err != nil {
			panic(err)
			return
		}
	})

	return cfg
}
