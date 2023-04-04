package sys

import (
	"context"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	isysLog "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	repositoryimpl "github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl/sys"
	"github.com/jettjia/go-ddd-demo/types"
)

// SysLogSvc sys_log_svc
//
//go:generate mockgen --source ./sys_log_svc.go --destination ./mock/mock_sys_log_svc.go --package mock
type SysLogSvc interface {
	CreateSysLog(ctx context.Context, sysLogEn *entity.SysLog) (id uint64, err error)                                                                                             // 创建
	DeleteSysLog(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                                        // 删除
	UpdateSysLog(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                                        // 修改
	FindSysLogById(ctx context.Context, id uint64) (sysLogEn *entity.SysLog, err error)                                                                                           // 查看byId
	FindSysLogByQuery(ctx context.Context, queries []*types.Query) (sysLogEn *entity.SysLog, err error)                                                                           // 查看byQuery
	FindSysLogAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysLog, err error)                                                                              // 所有
	FindSysLogPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysLog, pageData *types.PageData, err error) // 列表
	ExecSql(sql string) error                                                                                                                                                     // 执行sql
	FindOneExecSql(sql string) (sysLogEn *entity.SysLog, err error)                                                                                                               // 原生sql查询单个
	FindManyExecSql(sql string) (entries []*entity.SysLog, err error)                                                                                                             // 原生sql查询多个
}

type sysLogSvc struct {
	sysLogImpl isysLog.ISysLogRepo
}

func NewSysLogSvc() *sysLogSvc {
	return &sysLogSvc{
		sysLogImpl: repositoryimpl.NewSysLogImpl(),
	}
}

func (a *sysLogSvc) CreateSysLog(ctx context.Context, sysLogEn *entity.SysLog) (id uint64, err error) {
	return a.sysLogImpl.Create(ctx, sysLogEn)
}

func (a *sysLogSvc) DeleteSysLog(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	return a.sysLogImpl.Delete(ctx, sysLogEn)
}

func (a *sysLogSvc) UpdateSysLog(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	return a.sysLogImpl.Update(ctx, sysLogEn)
}

func (a *sysLogSvc) FindSysLogById(ctx context.Context, id uint64) (sysLogEn *entity.SysLog, err error) {
	return a.sysLogImpl.FindById(ctx, id)
}

func (a *sysLogSvc) FindSysLogByQuery(ctx context.Context, queries []*types.Query) (sysLogEn *entity.SysLog, err error) {
	return a.sysLogImpl.FindByQuery(ctx, queries)
}

func (a *sysLogSvc) FindSysLogAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysLog, err error) {
	return a.sysLogImpl.FindAll(ctx, queries)
}

func (a *sysLogSvc) FindSysLogPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysLog, pageData *types.PageData, err error) {
	return a.sysLogImpl.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *sysLogSvc) ExecSql(sql string) error {
	return a.sysLogImpl.ExecSql(sql)
}

func (a *sysLogSvc) FindOneExecSql(sql string) (sysLogEn *entity.SysLog, err error) {
	return a.sysLogImpl.FindOneExecSql(sql)
}

func (a *sysLogSvc) FindManyExecSql(sql string) (entries []*entity.SysLog, err error) {
	return a.sysLogImpl.FindManyExecSql(sql)
}
