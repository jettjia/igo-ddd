package aggregate

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	isysMenu "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	repositoryimpl "github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl/sys"
	"github.com/jettjia/go-ddd-demo/types"
)

// SysMenuAgg sys_menu_agg
//
//go:generate mockgen --source ./sys_menu_agg.go --destination ./mock/mock_sys_menu_agg.go --package mock
type SysMenuAgg interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error)                                                                                            // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                       // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                       // 修改
	FindSysMenuById(ctx context.Context, id uint64) (sysMenuEn *entity.SysMenu, err error)                                                                                          // 查看byId
	FindSysMenuByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error)                                                                          // 查看byQuery
	FindSysMenuAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error)                                                                              // 所有
	FindSysMenuPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysMenu, pageData *types.PageData, err error) // 列表
	ExecSql(sql string) error                                                                                                                                                       // 执行sql
	FindOneExecSql(sql string) (sysMenuEn *entity.SysMenu, err error)                                                                                                               // 原生sql查询单个
	FindManyExecSql(sql string) (entries []*entity.SysMenu, err error)                                                                                                              // 原生sql查询多个
}

type sysMenuAgg struct {
	sysMenuImpl isysMenu.ISysMenuRepo
}

func NewSysMenuAgg() *sysMenuAgg {
	return &sysMenuAgg{
		sysMenuImpl: repositoryimpl.NewSysMenuImpl(),
	}
}

func (a *sysMenuAgg) CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error) {
	return a.sysMenuImpl.Create(ctx, sysMenuEn)
}

func (a *sysMenuAgg) DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuImpl.Delete(ctx, sysMenuEn)
}

func (a *sysMenuAgg) UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuImpl.Update(ctx, sysMenuEn)
}

func (a *sysMenuAgg) FindSysMenuById(ctx context.Context, id uint64) (sysMenuEn *entity.SysMenu, err error) {
	sysMenuEn, err = a.sysMenuImpl.FindById(ctx, id)
	if err != nil {
		err = gerror.NewCode(responseutil.CommNotFound, err.Error())
		return
	}

	return
}

func (a *sysMenuAgg) FindSysMenuByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error) {
	sysMenuEn, err = a.sysMenuImpl.FindByQuery(ctx, queries)
	if err != nil {
		err = gerror.NewCode(responseutil.CommNotFound, err.Error())
		return
	}
	return
}

func (a *sysMenuAgg) FindSysMenuAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error) {
	return a.sysMenuImpl.FindAll(ctx, queries)
}

func (a *sysMenuAgg) FindSysMenuPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysMenu, pageData *types.PageData, err error) {
	return a.sysMenuImpl.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *sysMenuAgg) ExecSql(sql string) error {
	sql = `update sys_menu set menu_name = "a_menu"`
	return a.sysMenuImpl.ExecSql(sql)
}

func (a *sysMenuAgg) FindOneExecSql(sql string) (sysMenuEn *entity.SysMenu, err error) {
	//sql = `select * from sys_menu where id = 1`
	return a.sysMenuImpl.FindOneExecSql(sql)
}

func (a *sysMenuAgg) FindManyExecSql(sql string) (entries []*entity.SysMenu, err error) {
	//sql = `select * from sys_menu`
	return a.sysMenuImpl.FindManyExecSql(sql)
}
