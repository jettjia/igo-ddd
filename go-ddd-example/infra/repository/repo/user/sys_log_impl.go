package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/data"
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entitySysLog "github.com/jettjia/ddddemo/domain/entity/user"
	irepositorySysLog "github.com/jettjia/ddddemo/domain/irepository/user"
	"github.com/jettjia/ddddemo/infra/pkg/idata"
	converterSysLog "github.com/jettjia/ddddemo/infra/repository/converter/user"
	poSysLog "github.com/jettjia/ddddemo/infra/repository/po/user"
)

var _ irepositorySysLog.ISysLogRepo = (*SysLog)(nil)

type SysLog struct {
	data *data.Data
}

func NewSysLogImpl() *SysLog {
	return &SysLog{
		data: idata.NewDataOptionCli(),
	}
}

func (r *SysLog) Create(ctx context.Context, sysMenuEn *entitySysLog.SysLog) (ulid string, err error) {
	sysMenuPo := converterSysLog.E2PSysLogAdd(sysMenuEn)
	if err = r.data.DB(ctx).Create(&sysMenuPo).Error; err != nil {
		return
	}

	return sysMenuPo.Ulid, nil
}

func (r *SysLog) Delete(ctx context.Context, sysMenuEn *entitySysLog.SysLog) (err error) {
	sysMenuPo := converterSysLog.E2PSysLogDel(sysMenuEn)

	return r.data.DB(ctx).Model(&poSysLog.SysLog{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysLog) Update(ctx context.Context, sysMenuEn *entitySysLog.SysLog) (err error) {
	sysMenuPo := converterSysLog.E2PSysLogUpdate(sysMenuEn)

	return r.data.DB(ctx).Model(&poSysLog.SysLog{}).Where("ulid = ? ", sysMenuEn.Ulid).Updates(sysMenuPo).Error
}

func (r *SysLog) FindById(ctx context.Context, ulid string, selectColumn ...string) (sysMenuEn *entitySysLog.SysLog, err error) {
	var sysMenuPo poSysLog.SysLog
	if err = r.data.DB(ctx).Select(selectColumn).Limit(1).Find(&sysMenuPo, "ulid = ? ", ulid).Error; err != nil {
		return
	}

	sysMenuEn = converterSysLog.P2ESysLog(&sysMenuPo)

	return
}

func (r *SysLog) FindByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entitySysLog.SysLog, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	var sysMenuPo poSysLog.SysLog
	if err = r.data.DB(ctx).Model(&poSysLog.SysLog{}).Limit(1).Where(whereStr, values...).Find(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterSysLog.P2ESysLog(&sysMenuPo)

	return
}

func (r *SysLog) FindAll(ctx context.Context, queries []*builder.Query, selectArgs ...[]string) (entries []*entitySysLog.SysLog, err error) {
	whereStr, values, err := builder.GormBuildWhere(queries)
	if err != nil {
		return
	}

	selectField := builder.BuildSelectVariable(selectArgs...)

	sysMenuPos := make([]*poSysLog.SysLog, 0)
	if err = r.data.DB(ctx).Model(&poSysLog.SysLog{}).Select(selectField).Where(whereStr, values...).Order("ulid desc").Find(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterSysLog.P2ESysLogs(sysMenuPos)

	return
}

func (r *SysLog) FindPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData, selectArgs ...[]string) (entries []*entitySysLog.SysLog, rspPag *builder.PageData, err error) {
	var (
		total int64
	)
	sysMenuPos := make([]*poSysLog.SysLog, 0)

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

	dbQuery := r.data.DB(ctx).Model(&poSysLog.SysLog{}).Where(whereStr, values...)

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

	entries = converterSysLog.P2ESysLogs(sysMenuPos)

	return
}

func (r *SysLog) ExecSql(ctx context.Context, sql string) error {
	return r.data.DB(ctx).Exec(sql).Error
}

func (r *SysLog) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entitySysLog.SysLog, err error) {
	var sysMenuPo poSysLog.SysLog
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPo).Error; err != nil {
		return
	}

	sysMenuEn = converterSysLog.P2ESysLog(&sysMenuPo)

	return
}

func (r *SysLog) FindManyExecSql(ctx context.Context, sql string) (entries []*entitySysLog.SysLog, err error) {
	sysMenuPos := make([]*poSysLog.SysLog, 0)
	if err = r.data.DB(ctx).Raw(sql).Scan(&sysMenuPos).Error; err != nil {
		return
	}

	entries = converterSysLog.P2ESysLogs(sysMenuPos)

	return
}
