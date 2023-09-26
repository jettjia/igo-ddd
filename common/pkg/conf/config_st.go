package conf

type Config struct {
	Server  ServerConf
	Gserver GServerConf
	Mysql   MysqlConf
	Log     LogConf
	Nsq     NsqConf
	Otel    OtelConf
	Redis   RedisConf
}

// ServerConf http服务
type ServerConf struct {
	Lang              string `yaml:"lang"`
	PublicPort        int    `yaml:"public_port"`
	PrivatePort       int    `yaml:"private_port"`
	PublicMetricPort  int    `yaml:"public_metric_port"`
	PrivateMetricPort int    `yaml:"private_metric_port"`
	ServerName        string `yaml:"server_name"`
	Mode              string `yaml:"mode"`
	Dev               bool   `yaml:"dev"`
	EnableEvent       bool   `yaml:"enable_event"`
	EnableJob         bool   `yaml:"enable_job"`
	EnableGrpc        bool   `yaml:"enable_grpc"`
}

// GServerConf grpc服务
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
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	LogMode         int    `yaml:"log_mode"`
	SlowThreshold   int    `yaml:"slow_threshold"`
}

// RedisConf redis
type RedisConf struct {
	RedisType  string `yaml:"redis_type"` // redis使用模式:alone, sentinel,cluster
	Addr       string `yaml:"addr"`
	Password   string `yaml:"password"`
	MasterName string `yaml:"master_name"`
	PoolSize   int    `yaml:"pool_size"`
}

// LogConf log
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

// OtelConf otel
type OtelConf struct {
	Enable         bool   `yaml:"enable"`
	ExportEndpoint string `yaml:"export_endpoint"`
}
