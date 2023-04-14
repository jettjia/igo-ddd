package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"sync"
)

// ServerConf 服务端口，名称，mode
type ServerConf struct {
	Lang        string `yaml:"lang"`
	PublicPort  int    `yaml:"public_port"`
	PrivatePort int    `yaml:"private_port"`
	ServerName  string `yaml:"server_name"`
	Mode        string `yaml:"mode"`
}

// GServerConf grpc服务端口
type GServerConf struct {
	Host            string `yaml:"host"`
	PublicPort      int    `yaml:"public_port"`
	MaxMsgSize      int    `yaml:"max_msg_size"`
	ClientGoodsHost string `yaml:"client_goods_host"`
	ClientGoodsPort int    `yaml:"client_goods_port"`
}

// MysqlConf mysql
type MysqlConf struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbHost          string `yaml:"db_host"`
	DbPort          int    `yaml:"db_port"`
	DbName          string `yaml:"db_name"`
	Charset         string `yaml:"charset"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	LogMode         int    `yaml:"log_mode"`
	SlowThreshold   int    `yaml:"slow_threshold"`
}

// LogConf 日志
type LogConf struct {
	LogFileDir string `yaml:"log_file_dir"` // 日志目录
	AppName    string `yaml:"app_name"`     //日志名称
	MaxSize    int    `yaml:"max_size"`     //文件多大开始切分
	MaxBackups int    `yaml:"max_backups"`  //保留文件个数
	MaxAge     int    `yaml:"max_age"`      //文件保留最大实际
	LogLevel   string `yaml:"log_level"`    // 日志级别
	LogOut     string `yaml:"log_out"`      // 日志输出位置
}

// NsqConf nsq
type NsqConf struct {
	NsqProducerHost  string `yaml:"nsq_producer_host"`
	NsqProducerPort  string `yaml:"nsq_producer_port"`
	NsqSubscribeHost string `yaml:"nsq_subscribe_host"`
	NsqSubscribePort string `yaml:"nsq_subscribe_port"`
}

type Config struct {
	Server  ServerConf
	Gserver GServerConf
	Mysql   MysqlConf
	Log     LogConf
	Nsq     NsqConf
}

func NewConfig(env string) (conf *Config) {
	conf = initConfig(env)
	return
}

var (
	configOnce sync.Once
	config     *Config
)

// InitConfig 读取配置
func initConfig(env string) *Config {
	if len(env) == 0 {
		env = "debug"
	}
	configOnce.Do(func() {
		fileBytes, err := ioutil.ReadFile("./manifest/config/config-" + env + ".yaml")
		if err != nil {
			panic(fmt.Sprintf("load config.yaml failed: %v", err))
		}

		err = yaml.Unmarshal(fileBytes, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
	})

	return config
}
