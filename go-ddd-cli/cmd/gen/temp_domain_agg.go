package gen

const TempDomainAgg = `
package user

import (
	"context"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
	repoUser "xdata/xtext/zsvc-example/infra/repository/repo/user"
)

// SysMenuAgg sys_menu_agg
//
//go:generate mockgen --source ./sys_menu_agg.go --destination ./mock/mock_sys_menu_agg.go --package mock
type SysMenuAgg interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error) // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)              // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)              // 修改
}

type SysMenu struct {
	sysMenuRepo *repoUser.SysMenu
}

func NewSysMenuAgg() *SysMenu {
	return &SysMenu{
		sysMenuRepo: repoUser.NewSysMenuImpl(),
	}
}

func (a *SysMenu) CreateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error) {
	return a.sysMenuRepo.Create(ctx, sysMenuEn)
}

func (a *SysMenu) DeleteSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error) {
	return a.sysMenuRepo.Delete(ctx, sysMenuEn)
}

func (a *SysMenu) UpdateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error) {
	return a.sysMenuRepo.Update(ctx, sysMenuEn)
}

`
