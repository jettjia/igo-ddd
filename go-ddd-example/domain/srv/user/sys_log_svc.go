package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
	repoUser "github.com/jettjia/ddddemo/infra/repository/repo/user"
)

// SysLogSvc sys_log_svc
//
//go:generate mockgen --source ./sys_log_svc.go --destination ./mock/mock_sys_log_svc.go --package mock
type SysLogSvc interface {
	CreateSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (ulid string, err error)                                                                                                   // 创建
	DeleteSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (err error)                                                                                                                // 删除
	UpdateSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (err error)                                                                                                                // 修改
	FindSysLogById(ctx context.Context, ulid string) (sysLogEn *entityUser.SysLog, err error)                                                                                                 // 查看byId
	FindSysLogByQuery(ctx context.Context, queries []*builder.Query) (sysLogEn *entityUser.SysLog, err error)                                                                                 // 查看byQuery
	FindSysLogAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysLog, err error)                                                                                    // 所有
	FindSysLogPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysLog, pageData *builder.PageData, err error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                            // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysLogEn *entityUser.SysLog, err error)                                                                                                  // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysLog, err error)                                                                                                // 原生sql查询多个
}

type SysLog struct {
	sysLogRepo *repoUser.SysLog
}

func NewSysLogSvc() *SysLog {
	return &SysLog{
		sysLogRepo: repoUser.NewSysLogImpl(),
	}
}

func (a *SysLog) CreateSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (ulid string, err error) {
	return a.sysLogRepo.Create(ctx, sysLogEn)
}

func (a *SysLog) DeleteSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (err error) {
	return a.sysLogRepo.Delete(ctx, sysLogEn)
}

func (a *SysLog) UpdateSysLog(ctx context.Context, sysLogEn *entityUser.SysLog) (err error) {
	return a.sysLogRepo.Update(ctx, sysLogEn)
}

func (a *SysLog) FindSysLogById(ctx context.Context, ulid string) (sysLogEn *entityUser.SysLog, err error) {
	return a.sysLogRepo.FindById(ctx, ulid)
}

func (a *SysLog) FindSysLogByQuery(ctx context.Context, queries []*builder.Query) (sysLogEn *entityUser.SysLog, err error) {
	return a.sysLogRepo.FindByQuery(ctx, queries)
}

func (a *SysLog) FindSysLogAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysLog, err error) {
	return a.sysLogRepo.FindAll(ctx, queries)
}

func (a *SysLog) FindSysLogPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysLog, pageData *builder.PageData, err error) {
	return a.sysLogRepo.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *SysLog) ExecSql(ctx context.Context, sql string) error {
	return a.sysLogRepo.ExecSql(ctx, sql)
}

func (a *SysLog) FindOneExecSql(ctx context.Context, sql string) (sysLogEn *entityUser.SysLog, err error) {
	return a.sysLogRepo.FindOneExecSql(ctx, sql)
}

func (a *SysLog) FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysLog, err error) {
	return a.sysLogRepo.FindManyExecSql(ctx, sql)
}
