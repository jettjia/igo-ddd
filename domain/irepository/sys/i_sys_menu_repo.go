package irepository

import (
	"context"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	"github.com/jettjia/go-ddd-demo/types"
)

// ISysMenuRepo sys_menu
//
//go:generate mockgen --source ./i_sys_menu_repo.go --destination ./mock/mock_i_sys_menu_repo.go --package mock
type ISysMenuRepo interface {
	Create(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error)                                                                       // 创建
	Delete(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                  // 删除
	Update(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                  // 修改
	FindById(ctx context.Context, id uint64) (sysMenuEn *entity.SysMenu, err error)                                                                     // 查看byId
	FindByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error)                                                     // 查看byQuery
	FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error)                                                         // 所有
	FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) ([]*entity.SysMenu, *types.PageData, error) // 列表
	ExecSql(sql string) error                                                                                                                           // 执行sql
	FindOneExecSql(sql string) (sysMenuEn *entity.SysMenu, err error)                                                                                   // 原生sql查询单个
	FindManyExecSql(sql string) (entries []*entity.SysMenu, err error)                                                                                  // 原生sql查询多个
}
