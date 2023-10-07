package irepository

import (
	"context"

	"jettjia/go-ddd-demo-multi-common/types"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

// ISysMenuRepo sys_menu
//
//go:generate mockgen --source ./i_sys_menu_repo.go --destination ./mock/mock_i_sys_menu_repo.go --package mock
type ISysMenuRepo interface {
	Create(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error)                                                                                             // 创建
	Delete(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                          // 删除
	Update(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                          // 修改
	FindById(ctx context.Context, ulid string, selectColumn ...string) (sysMenuEn *entity.SysMenu, err error)                                                                   // 查看byId
	FindByQuery(ctx context.Context, queries []*types.Query, selectColumn ...string) (sysMenuEn *entity.SysMenu, err error)                                                     // 查看byQuery
	FindAll(ctx context.Context, queries []*types.Query, selectColumn ...string) (entries []*entity.SysMenu, err error)                                                         // 所有
	FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData, selectColumn ...string) ([]*entity.SysMenu, *types.PageData, error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                              // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entity.SysMenu, err error)                                                                                      // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysMenu, err error)                                                                                     // 原生sql查询多个
}
