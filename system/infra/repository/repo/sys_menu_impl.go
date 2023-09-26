package repo

import (
	"context"

	"jettjia/go-ddd-demo-multi-common/pkg/data"
	"jettjia/go-ddd-demo-multi-common/pkg/util"
	"jettjia/go-ddd-demo-multi-common/types"

	"jettjia/go-ddd-demo-multi-system/domain/entity"
	isysMenu "jettjia/go-ddd-demo-multi-system/domain/irepository"
	"jettjia/go-ddd-demo-multi-system/infra/repository/converter"
	"jettjia/go-ddd-demo-multi-system/infra/repository/po"
)

var _ isysMenu.ISysMenuRepo = (*SysMenu)(nil)

type SysMenu struct {
	data *data.Data
}

func NewSysMenuImpl(data *data.Data) *SysMenu {
	return &SysMenu{
		data: data,
	}
}

func (r *SysMenu) Create(ctx context.Context, sysMenuEn *entity.SysMenu) (ulid string, err error) {
	sysMenuPo := converter.E2PSysMenuAdd(sysMenuEn)
	err = r.data.DB(ctx).Create(&sysMenuPo).Error
	if err != nil {
		return
	}

	return sysMenuPo.Ulid, nil
}

func (r SysMenu) Delete(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	sysMenuPo := converter.E2PSysMenuDel(sysMenuEn)

	return r.data.DB(ctx).Model(&po.SysMenu{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysMenu) Update(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	sysMenuPo := converter.E2PSysMenuUpdate(sysMenuEn)

	return r.data.DB(ctx).Model(&po.SysMenu{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysMenu) FindById(ctx context.Context, ulid string) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	err = r.data.DB(ctx).Limit(1).Find(&sysMenuPo, "ulid = ? ", ulid).Error
	if err != nil {
		return sysMenuEn, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	condition := types.GenerateQueryCondition(queries)
	err = r.data.DB(ctx).Where(condition).Limit(1).Find(&sysMenuPo).Error
	if err != nil {
		return sysMenuEn, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error) {
	sysMenuPos := make([]*po.SysMenu, 0)
	condition := types.GenerateQueryCondition(queries)

	err = r.data.DB(ctx).Find(&sysMenuPos, condition).Error
	if err != nil {
		return entries, err
	}
	entries = converter.P2ESysMenus(sysMenuPos)

	return
}

func (r *SysMenu) FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) ([]*entity.SysMenu, *types.PageData, error) {
	var condition string
	var total int64
	sysMenuPos := make([]*po.SysMenu, 0)
	var rspPag types.PageData
	var entries []*entity.SysMenu

	condition = types.GenerateQueryCondition(queries)

	dbQuery := r.data.DB(ctx).Model(&po.SysMenu{}).Where(condition)

	err := dbQuery.Count(&total).Error
	if err != nil {
		return entries, &rspPag, err
	}

	if total != 0 {
		err = dbQuery.
			Select("sys_menu.*").
			Order(reqSort.Sort + " " + reqSort.Direction).
			Scopes(types.Paginate(reqPage.PageNum, reqPage.PageSize)).
			Find(&sysMenuPos).Error

		if err != nil {
			return entries, &rspPag, err
		}
	}

	rspPag.PageNum = reqPage.PageNum
	rspPag.PageSize = reqPage.PageSize
	rspPag.TotalNumber = total
	rspPag.TotalPage = util.CeilPageNum(total, reqPage.PageSize)

	entries = converter.P2ESysMenus(sysMenuPos)

	return entries, &rspPag, nil
}

func (r *SysMenu) ExecSql(ctx context.Context, sql string) error {
	return r.data.DB(ctx).Exec(sql).Error
}

func (r *SysMenu) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPo).Error
	if err != nil {
		return nil, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysMenu, err error) {
	sysMenuPos := make([]*po.SysMenu, 0)
	err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPos).Error
	if err != nil {
		return nil, err
	}
	entries = converter.P2ESysMenus(sysMenuPos)

	return
}
