package gen

const TempInfraImpl = `
package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/data"
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
	irepositoryUser "xdata/xtext/zsvc-example/domain/irepository/user"
	"xdata/xtext/zsvc-example/infra/pkg/idata"
	converterUser "xdata/xtext/zsvc-example/infra/repository/converter/user"
	poUser "xdata/xtext/zsvc-example/infra/repository/po/user"
)

var _ irepositoryUser.ISysMenuRepo = (*SysMenu)(nil)

type SysMenu struct {
	data *data.Data
}

func NewSysMenuImpl() *SysMenu {
	return &SysMenu{
		data: idata.NewDataOptionCli(),
	}
}

func (r *SysMenu) Create(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error) {
	sysMenuPo := converterUser.E2PSysMenuAdd(sysMenuEn)
	if err = r.data.DB(ctx).Create(&sysMenuPo).Error; err != nil {
		return
	}

	return sysMenuPo.Ulid, nil
}

func (r *SysMenu) Delete(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error) {
	sysMenuPo := converterUser.E2PSysMenuDel(sysMenuEn)

	return r.data.DB(ctx).Model(&poUser.SysMenu{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysMenu) Update(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error) {
	sysMenuPo := converterUser.E2PSysMenuUpdate(sysMenuEn)

	return r.data.DB(ctx).Model(&poUser.SysMenu{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysMenu) FindById(ctx context.Context, ulid string, selectColumn ...string) (sysMenuEn *entityUser.SysMenu, err error) {
	var sysMenuPo poUser.SysMenu
	if err = r.data.DB(ctx).Select(selectColumn).Limit(1).Find(&sysMenuPo, "ulid = ? ", ulid).Error; err != nil {
		return
	}

	sysMenuEn = converterUser.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	var sysMenuPo poUser.SysMenu
	if err = r.data.DB(ctx).Model(&poUser.SysMenu{}).Limit(1).Where(whereStr, values...).Find(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterUser.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindAll(ctx context.Context, queries []*builder.Query, selectArgs ...[]string) (entries []*entityUser.SysMenu, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	selectField := builder.BuildSelectVariable(selectArgs...)

	sysMenuPos := make([]*poUser.SysMenu, 0)
	if err = r.data.DB(ctx).Model(&poUser.SysMenu{}).Select(selectField).Where(whereStr, values...).Order("ulid desc").Find(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterUser.P2ESysMenus(sysMenuPos)

	return
}

func (r *SysMenu) DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	err = r.data.DB(ctx).Unscoped().Where(whereStr, values...).Delete(&poUser.SysMenu{}).Error

	return
}

func (r *SysMenu) DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	sysMenuPo := converterUser.E2PSysMenuDel(sysMenuEn)

	err = r.data.DB(ctx).Where(whereStr, values...).Updates(sysMenuPo).Error

	return
}

func (r *SysMenu) FindPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData, selectArgs ...[]string) (entries []*entityUser.SysMenu, rspPag *builder.PageData, err error) {
	var (
		total int64
	)
	sysMenuPos := make([]*poUser.SysMenu, 0)

	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	// default reqSort
	if reqSort == nil {
		reqSort = &builder.SortData{Sort: "ulid", Direction: "desc"}
	}
	// default reqPage
	if reqPage == nil {
		reqPage = &builder.PageData{PageNum: 1, PageSize: 10}
	}
	// select
	selectField := builder.BuildSelectVariable(selectArgs...)

	dbQuery := r.data.DB(ctx).Model(&poUser.SysMenu{}).Where(whereStr, values...)

	if err = dbQuery.Count(&total).Error; err != nil {
		return
	}

	rspPag = &builder.PageData{
		PageNum:     reqPage.PageNum,
		PageSize:    reqPage.PageSize,
		TotalNumber: total,
		TotalPage:   builder.CeilPageNum(total, reqPage.PageSize),
	}

	if total == 0 {
		return
	}

	err = dbQuery.
		Select(selectField).
		Order(reqSort.Sort + " " + reqSort.Direction).
		Scopes(builder.GormPaginate(reqPage.PageNum, reqPage.PageSize)).
		Find(&sysMenuPos).
		Error

	if err != nil {
		return
	}

	entries = converterUser.P2ESysMenus(sysMenuPos)

	return
}

func (r *SysMenu) ExecSql(ctx context.Context, sql string) error {
	return r.data.DB(ctx).Exec(sql).Error
}

func (r *SysMenu) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error) {
	var sysMenuPo poUser.SysMenu
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterUser.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error) {
	sysMenuPos := make([]*poUser.SysMenu, 0)
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterUser.P2ESysMenus(sysMenuPos)

	return
}
`
