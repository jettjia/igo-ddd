package config

import (
	"flag"
	"os"
	"sync"

	"github.com/jettjia/igo-pkg/pkg/conf"
)

var (
	configOnce sync.Once
	cfg        *conf.Config
)

func NewConfig() (conf *conf.Config) {
	configOnce.Do(func() {
		cfg = initConfig()
	})

	return cfg
}

// InitConfig read config
func initConfig() *conf.Config {
	// dev: load location config path
	file := "./manifest/config/config.yaml"

	// product: load k8s config path
	if os.Getenv("env") == "release" {
		file = "/sysvol/conf/public-center.yaml"
	}

	// product: load docker config path
	if os.Getenv("env") == "docker" {
		file = "./manifest/config/config-docker.yaml"
	}

	// load config.yaml
	cfg = &conf.Config{}
	if err := flag.Set("conf", file); err != nil {
		panic(err)
	}
	if err := conf.ParseYaml(cfg); err != nil {
		panic(err)
	}

	if err := conf.ParseYaml(cfg); err != nil {
		panic(err)
	}

	return cfg
}
