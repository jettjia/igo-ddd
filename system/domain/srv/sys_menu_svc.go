package srv

import (
	"context"

	"jettjia/go-ddd-demo-multi-common/types"

	"jettjia/go-ddd-demo-multi-system/domain/entity"
	"jettjia/go-ddd-demo-multi-system/infra/repository/repo"
)

// SysMenuSvc sys_menu_svc
//
//go:generate mockgen --source ./sys_menu_svc.go --destination ./mock/mock_sys_menu_svc.go --package mock
type SysMenuSvc interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error)                                                                                          // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                       // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)                                                                                                       // 修改
	FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entity.SysMenu, err error)                                                                                        // 查看byId
	FindSysMenuByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error)                                                                          // 查看byQuery
	FindSysMenuAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error)                                                                              // 所有
	FindSysMenuPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysMenu, pageData *types.PageData, err error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                  // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entity.SysMenu, err error)                                                                                          // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysMenu, err error)                                                                                         // 原生sql查询多个
}

type SysMenu struct {
	sysMenuRepo *repo.SysMenu
}

func NewSysMenuSvc(sysMenuRepo *repo.SysMenu) *SysMenu {
	return &SysMenu{
		sysMenuRepo: sysMenuRepo,
	}
}

func (a *SysMenu) CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error) {
	return a.sysMenuRepo.Create(ctx, sysMenuEn)
}

func (a *SysMenu) DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuRepo.Delete(ctx, sysMenuEn)
}

func (a *SysMenu) UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	return a.sysMenuRepo.Update(ctx, sysMenuEn)
}

func (a *SysMenu) FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entity.SysMenu, err error) {
	return a.sysMenuRepo.FindById(ctx, ulid)
}

func (a *SysMenu) FindSysMenuByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error) {
	return a.sysMenuRepo.FindByQuery(ctx, queries)
}

func (a *SysMenu) FindSysMenuAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error) {
	return a.sysMenuRepo.FindAll(ctx, queries)
}

func (a *SysMenu) FindSysMenuPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) (entries []*entity.SysMenu, pageData *types.PageData, err error) {
	return a.sysMenuRepo.FindPage(ctx, queries, reqPage, reqSort, "ulid")
}

func (a *SysMenu) ExecSql(ctx context.Context, sql string) error {
	return a.sysMenuRepo.ExecSql(ctx, sql)
}

func (a *SysMenu) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entity.SysMenu, err error) {
	return a.sysMenuRepo.FindOneExecSql(ctx, sql)
}

func (a *SysMenu) FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysMenu, err error) {
	return a.sysMenuRepo.FindManyExecSql(ctx, sql)
}
