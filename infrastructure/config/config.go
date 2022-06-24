package config

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

// ServerConf 服务端口，名称，mode
type ServerConf struct {
	Address    string `yaml:"address"`
	ServerName string `yaml:"serverName"`
	Mode       string `yaml:"mode"`
}

// MysqlConf mysql
type MysqlConf struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbHost   string `yaml:"dbHost"`
	DbPort   int    `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
}

// LogConf 日志
type LogConf struct {
	LogFileDir string `yaml:"logFileDir"`
	AppName    string `yaml:"appName"`
	MaxSize    int    `yaml:"maxSize"`    //文件多大开始切分
	MaxBackups int    `yaml:"maxBackups"` //保留文件个数
	MaxAge     int    `yaml:"maxAge"`     //文件保留最大实际
}

// NsqConf nsq
type NsqConf struct {
	NsqProducerHost  string `yaml:"nsq_producer_host"`
	NsqProducerPort  string `yaml:"nsq_producer_port"`
	NsqSubscribeHost string `yaml:"nsq_subscribe_host"`
	NsqSubscribePort string `yaml:"nsq_subscribe_port"`
}

type Config struct {
	Server ServerConf
	Mysql  MysqlConf
	Log    LogConf
	Nsq    NsqConf
}

func NewConfig() (conf *Config) {
	conf = initConfig()
	return
}

var (
	configOnce sync.Once
	config     *Config
)

// InitConfig 读取配置
func initConfig() *Config {
	configOnce.Do(func() {
		configFilePath := "/var/manifest/config/config.yaml"
		file, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			panic(fmt.Sprintf("load %v failed: %v", configFilePath, err))
		}

		err = yaml.Unmarshal(file, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
	})

	return config
}
