package sys

import (
	"context"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	"github.com/jettjia/go-ddd-demo/types"
)

// ISysLogRepo sys_log
//
//go:generate mockgen --source ./i_sys_log_repo.go --destination ./mock/mock_i_sys_log_repo.go --package mock
type ISysLogRepo interface {
	Create(ctx context.Context, sysLogEn *entity.SysLog) (id uint64, err error)                                                                        // 创建
	Delete(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                   // 删除
	Update(ctx context.Context, sysLogEn *entity.SysLog) (err error)                                                                                   // 修改
	FindById(ctx context.Context, id uint64) (sysLogEn *entity.SysLog, err error)                                                                      // 查看byId
	FindByQuery(ctx context.Context, queries []*types.Query) (sysLogEn *entity.SysLog, err error)                                                      // 查看byQuery
	FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysLog, err error)                                                         // 所有
	FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) ([]*entity.SysLog, *types.PageData, error) // 列表
	ExecSql(sql string) error                                                                                                                          // 执行sql
	FindOneExecSql(sql string) (sysLogEn *entity.SysLog, err error)                                                                                    // 原生sql查询单个
	FindManyExecSql(sql string) (entries []*entity.SysLog, err error)                                                                                  // 原生sql查询多个
}
