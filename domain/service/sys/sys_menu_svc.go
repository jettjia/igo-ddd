package sys

import (
	"context"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	isysMenu "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	repositoryimpl "github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl/sys"
	"github.com/jettjia/go-ddd-demo/types"
)

// SysMenuSvc sys_menu_svc
//
//go:generate mockgen --source ./sys_menu_svc.go --destination ./mock/mock_sys_menu_svc.go --package mock
type SysMenuSvc interface {
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

type sysMenuSvc struct {
	sysMenuImpl isysMenu.ISysMenuRepo
}

func NewSysMenuSvc() *sysMenuSvc {
	return &sysMenuSvc{
		sysMenuImpl: repositoryimpl.NewSysMenuImpl(),
	}
}

func (a *sysMenuSvc) CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error) {
	return a.sysMenuImpl.Create(ctx, sysMenuEn)
}

func (a *sysMenuSvc) DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuImpl.Delete(ctx, sysMenuEn)
}

func (a *sysMenuSvc) UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuImpl.Update(ctx, sysMenuEn)
}

func (a *sysMenuSvc) FindSysMenuById(ctx context.Context, id uint64) (sysMenuEn *entity.SysMenu, err error) {
	return a.sysMenuImpl.FindById(ctx, id)
}

func (a *sysMenuSvc) FindSysMenuByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error) {
	return a.sysMenuImpl.FindByQuery(ctx, queries)
}

func (a *sysMenuSvc) FindSysMenuAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error) {
	return a.sysMenuImpl.FindAll(ctx, queries)
}

func (a *sysMenuSvc) FindSysMenuPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysMenu, pageData *types.PageData, err error) {
	return a.sysMenuImpl.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *sysMenuSvc) ExecSql(sql string) error {
	sql = `update sys_menu set menu_name = "a_menu"`
	return a.sysMenuImpl.ExecSql(sql)
}

func (a *sysMenuSvc) FindOneExecSql(sql string) (sysMenuEn *entity.SysMenu, err error) {
	//sql = `select * from sys_menu where id = 1`
	return a.sysMenuImpl.FindOneExecSql(sql)
}

func (a *sysMenuSvc) FindManyExecSql(sql string) (entries []*entity.SysMenu, err error) {
	//sql = `select * from sys_menu`
	return a.sysMenuImpl.FindManyExecSql(sql)
}
