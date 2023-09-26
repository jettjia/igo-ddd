package aggregate

import (
	"context"
	"jettjia/go-ddd-demo-multi-common/pkg/data"

	"jettjia/go-ddd-demo-multi-system/domain/entity"
	"jettjia/go-ddd-demo-multi-system/infra/repository/repo"
)

// SysMenuAgg sys_menu_agg
//
//go:generate mockgen --source ./sys_menu_agg.go --destination ./mock/mock_sys_menu_agg.go --package mock
type SysMenuAgg interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error) // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)              // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error)              // 修改
}

type SysMenu struct {
	tx          data.Transaction
	sysMenuRepo *repo.SysMenu
	sysLogRepo  *repo.SysLog
}

func NewSysMenuAgg(tx data.Transaction, sysMenuRepo *repo.SysMenu, sysLogRepo *repo.SysLog) *SysMenu {
	return &SysMenu{
		tx:          tx,
		sysMenuRepo: sysMenuRepo,
		sysLogRepo:  sysLogRepo,
	}
}

func (a *SysMenu) CreateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error) {

	var (
		sysLogEn entity.SysLog
	)

	// 调用事务实例
	err = a.tx.ExecTx(ctx, func(ctx context.Context) error {
		ulid, err = a.sysMenuRepo.Create(ctx, sysMenuEn)
		if err != nil {
			return err
		}

		sysLogEn.CreatedBy = sysMenuEn.CreatedBy
		sysLogEn.Msg = "SysMenu.Create" + ",ulid" + ulid
		_, err = a.sysLogRepo.Create(ctx, &sysLogEn)

		return err
	})

	return
}

func (a *SysMenu) DeleteSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	if err = a.sysMenuRepo.Delete(ctx, sysMenuEn); err != nil {
		return err
	}
	var (
		sysLogEn entity.SysLog
	)

	sysLogEn.CreatedBy = sysMenuEn.DeletedBy
	sysLogEn.Msg = "SysMenu.Delete"
	_, err = a.sysLogRepo.Create(ctx, &sysLogEn)

	return
}

func (a *SysMenu) UpdateSysMenu(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	if err = a.sysMenuRepo.Update(ctx, sysMenuEn); err != nil {
		return err
	}

	var (
		sysLogEn entity.SysLog
	)

	sysLogEn.CreatedBy = sysMenuEn.UpdatedBy
	sysLogEn.Msg = "SysMenu.Update"
	_, err = a.sysLogRepo.Create(ctx, &sysLogEn)

	return
}
