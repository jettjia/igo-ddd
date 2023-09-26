package repo

import (
	"context"
	"jettjia/go-ddd-demo-multi-common/pkg/data"
	"jettjia/go-ddd-demo-multi-common/pkg/util"
	"jettjia/go-ddd-demo-multi-common/types"

	"jettjia/go-ddd-demo-multi-system/domain/entity"
	isysLog "jettjia/go-ddd-demo-multi-system/domain/irepository"
	"jettjia/go-ddd-demo-multi-system/infra/repository/converter"
	"jettjia/go-ddd-demo-multi-system/infra/repository/po"
)

var _ isysLog.ISysLogRepo = (*SysLog)(nil)

type SysLog struct {
	data *data.Data
}

func NewSysLogImpl(data *data.Data) *SysLog {
	return &SysLog{
		data: data,
	}
}

func (r *SysLog) Create(ctx context.Context, sysLogEn *entity.SysLog) (ulid string, err error) {
	sysLogPo := converter.E2PSysLogAdd(sysLogEn)
	err = r.data.DB(ctx).Create(&sysLogPo).Error
	if err != nil {
		return
	}

	return sysLogPo.Ulid, nil
}

func (r SysLog) Delete(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	sysLogPo := converter.E2PSysLogDel(sysLogEn)

	return r.data.DB(ctx).Model(&po.SysLog{}).Where("ulid = ? ", sysLogEn.Ulid).Updates(sysLogPo).Error
}

func (r *SysLog) Update(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	sysLogPo := converter.E2PSysLogUpdate(sysLogEn)

	return r.data.DB(ctx).Model(&po.SysLog{}).Where("ulid = ? ", sysLogEn.Ulid).Updates(sysLogPo).Error
}

func (r *SysLog) FindById(ctx context.Context, ulid string) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	err = r.data.DB(ctx).Limit(1).Find(&sysLogPo, "ulid = ? ", ulid).Error
	if err != nil {
		return sysLogEn, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindByQuery(ctx context.Context, queries []*types.Query) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	condition := types.GenerateQueryCondition(queries)
	err = r.data.DB(ctx).Where(condition).Find(1).First(&sysLogPo).Error
	if err != nil {
		return sysLogEn, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysLog, err error) {
	sysLogPos := make([]*po.SysLog, 0)
	condition := types.GenerateQueryCondition(queries)

	err = r.data.DB(ctx).Find(&sysLogPos, condition).Error
	if err != nil {
		return entries, err
	}
	entries = converter.P2ESysLogs(sysLogPos)

	return
}

func (r *SysLog) FindPage(ctx context.Context, queries []*types.Query, reqPage *types.PageData, reqSort *types.SortData) ([]*entity.SysLog, *types.PageData, error) {
	var condition string
	var total int64
	sysLogPos := make([]*po.SysLog, 0)
	var rspPag types.PageData
	var entries []*entity.SysLog

	condition = types.GenerateQueryCondition(queries)

	dbQuery := r.data.DB(ctx).Model(&po.SysLog{}).Where(condition)

	err := dbQuery.Count(&total).Error
	if err != nil {
		return entries, &rspPag, err
	}

	if total != 0 {
		err = dbQuery.
			Select("sys_log.*").
			Order(reqSort.Sort + " " + reqSort.Direction).
			Scopes(types.Paginate(reqPage.PageNum, reqPage.PageSize)).
			Find(&sysLogPos).Error
		if err != nil {
			return entries, &rspPag, err
		}
	}

	rspPag.PageNum = reqPage.PageNum
	rspPag.PageSize = reqPage.PageSize
	rspPag.TotalNumber = total
	rspPag.TotalPage = util.CeilPageNum(total, reqPage.PageSize)

	entries = converter.P2ESysLogs(sysLogPos)

	return entries, &rspPag, nil
}

func (r *SysLog) ExecSql(ctx context.Context, sql string) error {
	return r.data.DB(ctx).Exec(sql).Error
}

func (r *SysLog) FindOneExecSql(ctx context.Context, sql string) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	err = r.data.DB(ctx).Raw(sql).Scan(&sysLogPo).Error
	if err != nil {
		return nil, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindManyExecSql(ctx context.Context, sql string) (entries []*entity.SysLog, err error) {
	sysLogPos := make([]*po.SysLog, 0)
	err = r.data.DB(ctx).Raw(sql).Scan(&sysLogPos).Error
	if err != nil {
		return nil, err
	}
	entries = converter.P2ESysLogs(sysLogPos)

	return
}
