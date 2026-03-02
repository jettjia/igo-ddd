package gen

import (
	"jettjia/igo-ddd-cli/util"
	"os"

	"github.com/gogf/gf-cli/v2/library/mlog"
)

func Run(host, user, password, port, db, table, serverName, dbType string) {
	// 1. 获取表完整结构信息
	InitDB(host, port, user, password, db, dbType)

	genReq := GenInit(serverName, table)

	mlog.Print("auto gen code start...")
	// 2. 生成项目文件结构
	CreateDir(genReq)

	// 3. go.mod
	GenMod(genReq)

	// 4. 生成  domain
	// domain/entity
	// domain/irepository api & mock
	// domain/aggregate api & impl & mock
	// domain/service api & impl & mock
	GenDomainEntity(genReq)
	GenDomainIrepository(genReq)
	GenDomainAgg(genReq)
	GenDomainSvc(genReq)

	// 5.infrastructure
	// infrastructure/repository impl
	// infrastructure/po
	// infrastructure/converter
	GenModel(genReq)
	GenConverter(genReq)
	//GenInfraComm(genReq)
	GenInfraImpl(genReq)

	// 6. application
	// application/dto
	// application/assembler
	// application/service
	GenApplicationDto(genReq)
	GenApplicationAssembler(genReq)
	GenApplicationService(genReq)

	// 7. api
	// api/handler
	// api/router
	GenInterfacesHandler(genReq)
	GenInterfacesRouter(genReq)

	// 8. boot
	// boot/wire

	// 9. test
	// test/unit
	GenUt(genReq)

	// 自动化文档
	GenApiDoc(genReq)

	//  grpc
	GenGrpcCode(genReq)

	// 10.格式化代码
	util.GoFmt(genReq.BaseDir)

	mlog.Print("done!")
}

// CreateDir 创建需要的文件夹
func CreateDir(req GenReq) {
	os.MkdirAll(req.InfraDir, os.ModePerm)

	os.MkdirAll(req.ConverterDir, os.ModePerm)
	os.MkdirAll(req.PoDir, os.ModePerm)
	os.MkdirAll(req.RepositoryimplDir, os.ModePerm)

	os.MkdirAll(req.AggregateDir, os.ModePerm)
	os.MkdirAll(req.EntityDir, os.ModePerm)
	os.MkdirAll(req.IrepositoryDir, os.ModePerm)
	os.MkdirAll(req.ServiceDir, os.ModePerm)

	os.MkdirAll(req.DtoDir, os.ModePerm)
	os.MkdirAll(req.AssemblerDir, os.ModePerm)
	os.MkdirAll(req.AServiceDir, os.ModePerm)

	os.MkdirAll(req.HandlerDir, os.ModePerm)
	os.MkdirAll(req.RouterDir, os.ModePerm)

	//os.MkdirAll(req.TypesDir, os.ModePerm)

	os.MkdirAll(req.TestDir, os.ModePerm)

	os.MkdirAll(req.ApiDocDir, os.ModePerm)
}
