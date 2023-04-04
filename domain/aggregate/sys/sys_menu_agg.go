package aggregate

import (
	"context"
	repositoryimpl "github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl/sys"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	repoSys "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	"github.com/jettjia/go-ddd-demo/global"
)

// SysMenuAgg sys_menu_agg
//
//go:generate mockgen --source ./sys_menu_agg.go --destination ./mock/mock_sys_menu_agg.go --package mock
type SysMenuAgg interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error) // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)            // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)            // 修改
}

type sysMenuAgg struct {
	sysMenuRepo repoSys.ISysMenuRepo
	sysLogRepo  repoSys.ISysLogRepo
}

func NewSysMenuAgg() *sysMenuAgg {
	return &sysMenuAgg{
		sysMenuRepo: repositoryimpl.NewSysMenuImpl(),
		sysLogRepo:  repositoryimpl.NewSysLogImpl(),
	}
}

func (a *sysMenuAgg) CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error) {
	if id, err = a.sysMenuRepo.Create(ctx, sysMenuEn); err != nil {
		return 0, err
	}

	var (
		sysLogEn entity.SysLog
	)

	sysLogEn.CreatedBy = global.GCustomerInfo.Username
	sysLogEn.Msg = "SysMenu.Create"
	_, _ = a.sysLogRepo.Create(ctx, &sysLogEn)

	return
}

func (a *sysMenuAgg) DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	if err = a.sysMenuRepo.Delete(ctx, sysMenuEn); err != nil {
		return err
	}
	var (
		sysLogEn entity.SysLog
	)

	sysLogEn.CreatedBy = global.GCustomerInfo.Username
	sysLogEn.Msg = "SysMenu.Delete"
	_, _ = a.sysLogRepo.Create(ctx, &sysLogEn)

	return
}

func (a *sysMenuAgg) UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	if err = a.sysMenuRepo.Update(ctx, sysMenuEn); err != nil {
		return err
	}

	var (
		sysLogEn entity.SysLog
	)

	sysLogEn.CreatedBy = global.GCustomerInfo.Username
	sysLogEn.Msg = "SysMenu.Update"
	_, _ = a.sysLogRepo.Create(ctx, &sysLogEn)

	return
}
