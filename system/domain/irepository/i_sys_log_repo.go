package irepository

import (
	"context"

	"jettjia/go-ddd-demo-multi-common/types"

	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

// ISysLogRepo sys_log
//
//go:generate mockgen --source ./i_sys_log_repo.go --destination ./mock/mock_i_sys_log_repo.go --package mock
type ISysLogRepo interface {
	Create(ctx context.Context, sysLogEn *entity.SysLog) (ulid string, err error)                                                                                              // 创建
	Delete(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                                           // 删除
	Update(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                                           // 修改
	FindById(ctx context.Context, ulid string, selectColumn ...string) (sysLogEn *entity.SysLog, err error)                                                                    // 查看byId
	FindByQuery(ctx context.Context, queries []*types.Query, selectColumn ...string) (sysLogEn *entity.SysLog, err error)                                                      // 查看byQuery
	FindAll(ctx context.Context, queries []*types.Query, selectColumn ...string) (entries []*entity.SysLog, err error)                                                         // 所有
	FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData, selectColumn ...string) ([]*entity.SysLog, *types.PageData, error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                             // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysLogEn *entity.SysLog, err error)                                                                                       // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysLog, err error)                                                                                     // 原生sql查询多个
}
