package gen

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type TableColumn struct {
	Field   string `gorm:"column:Field"`   // 字段名称
	Type    string `gorm:"column:Type"`    // 字段类型
	Null    string `gorm:"column:Null"`    // 是否空
	Key     string `gorm:"column:Key"`     // 索引
	Default string `gorm:"column:Default"` // 默认值
	Extra   string `gorm:"column:Extra"`   // 扩展
	Comment string `gorm:"column:Comment"` // 备注
}

type GenReq struct {
	TableName         string // 表名
	SrvName           string // 服务名
	TableKey          string // 主键id
	BaseDir           string // 基础路径
	InfraDir          string // infra
	PoDir             string // infra/po
	ConverterDir      string // infra/converter
	RepositoryimplDir string // infra/repositoryimpl
	AggregateDir      string // domain/aggregate
	EntityDir         string // domain/entity
	IrepositoryDir    string // domain/irepository
	ServiceDir        string // domain/service
	AssemblerDir      string // application/assembler
	DtoDir            string // application/dto
	AServiceDir       string // application/service
	HandlerDir        string // api/http/handler
	RouterDir         string // api/http/router
	TestDir           string // test
	ApiDocDir         string // api doc
	GrpcDir           string // grpc
	TableColumns      []TableColumn
}

func GenInit(srvName, tableName string) GenReq {
	codePath := viper.Get("server.code_path").(string)
	baseDir := codePath + "/auto-code"

	serverName := viper.Get("server.server_name").(string)

	return GenReq{
		BaseDir:           baseDir,
		TableName:         tableName,
		TableKey:          viper.Get("mysql.key").(string),
		SrvName:           srvName,
		InfraDir:          baseDir + "/infra",
		ConverterDir:      baseDir + "/infra/repository/converter/" + serverName,
		PoDir:             baseDir + "/infra/repository/po/" + serverName,
		RepositoryimplDir: baseDir + "/infra/repository/repo/" + serverName,
		AggregateDir:      baseDir + "/domain/aggregate/" + serverName,
		EntityDir:         baseDir + "/domain/entity/" + serverName,
		IrepositoryDir:    baseDir + "/domain/irepository/" + serverName,
		ServiceDir:        baseDir + "/domain/srv/" + serverName,
		AssemblerDir:      baseDir + "/application/assembler/" + serverName,
		DtoDir:            baseDir + "/application/dto/" + serverName,
		AServiceDir:       baseDir + "/application/service/" + serverName,
		HandlerDir:        baseDir + "/api/http/handler/public/" + serverName,
		RouterDir:         baseDir + "/api/http/router/public",
		TestDir:           baseDir + "/test/",
		ApiDocDir:         baseDir + "/doc",
		GrpcDir:           baseDir + "/api/grpc",
		TableColumns:      GetTableCol(tableName),
	}
}
