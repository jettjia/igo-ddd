package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/data"
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entitySysUser "github.com/jettjia/ddddemo/domain/entity/user"
	irepositorySysUser "github.com/jettjia/ddddemo/domain/irepository/user"
	"github.com/jettjia/ddddemo/infra/pkg/idata"
	converterSysUser "github.com/jettjia/ddddemo/infra/repository/converter/user"
	poSysUser "github.com/jettjia/ddddemo/infra/repository/po/user"
)

var _ irepositorySysUser.ISysUserRepo = (*SysUser)(nil)

type SysUser struct {
	data *data.Data
}

func NewSysUserImpl() *SysUser {
	return &SysUser{
		data: idata.NewDataOptionCli(),
	}
}

func (r *SysUser) Create(ctx context.Context, sysMenuEn *entitySysUser.SysUser) (ulid string, err error) {
	sysMenuPo := converterSysUser.E2PSysUserAdd(sysMenuEn)
	if err = r.data.DB(ctx).Create(&sysMenuPo).Error; err != nil {
		return
	}

	return sysMenuPo.Ulid, nil
}

func (r *SysUser) Delete(ctx context.Context, sysMenuEn *entitySysUser.SysUser) (err error) {
	sysMenuPo := converterSysUser.E2PSysUserDel(sysMenuEn)

	return r.data.DB(ctx).Model(&poSysUser.SysUser{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysUser) Update(ctx context.Context, sysMenuEn *entitySysUser.SysUser) (err error) {
	sysMenuPo := converterSysUser.E2PSysUserUpdate(sysMenuEn)

	return r.data.DB(ctx).Model(&poSysUser.SysUser{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysUser) FindById(ctx context.Context, ulid string, selectColumn ...string) (sysMenuEn *entitySysUser.SysUser, err error) {
	var sysMenuPo poSysUser.SysUser
	if err = r.data.DB(ctx).Select(selectColumn).Limit(1).Find(&sysMenuPo, "ulid = ? ", ulid).Error; err != nil {
		return
	}

	sysMenuEn = converterSysUser.P2ESysUser(&sysMenuPo)

	return
}

func (r *SysUser) FindByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entitySysUser.SysUser, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	var sysMenuPo poSysUser.SysUser
	if err = r.data.DB(ctx).Model(&poSysUser.SysUser{}).Limit(1).Where(whereStr, values...).Find(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterSysUser.P2ESysUser(&sysMenuPo)

	return
}

func (r *SysUser) FindAll(ctx context.Context, queries []*builder.Query, selectArgs ...[]string) (entries []*entitySysUser.SysUser, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	selectField := builder.BuildSelectVariable(selectArgs...)

	sysMenuPos := make([]*poSysUser.SysUser, 0)
	if err = r.data.DB(ctx).Model(&poSysUser.SysUser{}).Select(selectField).Where(whereStr, values...).Order("ulid desc").Find(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterSysUser.P2ESysUsers(sysMenuPos)

	return
}

func (r *SysUser) FindPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData, selectArgs ...[]string) (entries []*entitySysUser.SysUser, rspPag *builder.PageData, err error) {
	var (
		total int64
	)
	sysMenuPos := make([]*poSysUser.SysUser, 0)

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

	dbQuery := r.data.DB(ctx).Model(&poSysUser.SysUser{}).Where(whereStr, values...)

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

	entries = converterSysUser.P2ESysUsers(sysMenuPos)

	return
}

func (r *SysUser) ExecSql(ctx context.Context, sql string) error {
	return r.data.DB(ctx).Exec(sql).Error
}

func (r *SysUser) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entitySysUser.SysUser, err error) {
	var sysMenuPo poSysUser.SysUser
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterSysUser.P2ESysUser(&sysMenuPo)

	return
}

func (r *SysUser) FindManyExecSql(ctx context.Context, sql string) (entries []*entitySysUser.SysUser, err error) {
	sysMenuPos := make([]*poSysUser.SysUser, 0)
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterSysUser.P2ESysUsers(sysMenuPos)

	return
}
