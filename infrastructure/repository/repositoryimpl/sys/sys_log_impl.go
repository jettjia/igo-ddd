package sys

import (
	"context"

	"gorm.io/gorm"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	isysLog "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
	converter "github.com/jettjia/go-ddd-demo/infrastructure/repository/converter/sys"
	po "github.com/jettjia/go-ddd-demo/infrastructure/repository/po/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl"
	"github.com/jettjia/go-ddd-demo/types"
)

var _ isysLog.ISysLogRepo = (*SysLog)(nil)

type SysLog struct {
	db *gorm.DB
}

func NewSysLogImpl() *SysLog {
	return &SysLog{
		db: global.GDB,
	}
}

func (r *SysLog) Create(ctx context.Context, sysLogEn *entity.SysLog) (id uint64, err error) {
	sysLogPo := converter.E2PSysLogAdd(sysLogEn)
	err = r.db.Create(&sysLogPo).Error
	if err != nil {
		return 0, err
	}

	return sysLogPo.Id, nil
}

func (r SysLog) Delete(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	sysLogPo := converter.E2PSysLogDel(sysLogEn)

	return r.db.Model(&po.SysLog{}).Where("id = ? ", sysLogEn.Id).Updates(sysLogPo).Error
}

func (r *SysLog) Update(ctx context.Context, sysLogEn *entity.SysLog) (err error) {
	sysLogPo := converter.E2PSysLogUpdate(sysLogEn)

	return r.db.Model(&po.SysLog{}).Where("id = ? ", sysLogEn.Id).Updates(sysLogPo).Error
}

func (r *SysLog) FindById(ctx context.Context, id uint64) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	err = r.db.Limit(1).First(&sysLogPo, "id = ? ", id).Error
	if err != nil {
		return sysLogEn, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindByQuery(ctx context.Context, queries []*types.Query) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	condition := repositoryimpl.GenerateQueryCondition(queries)
	err = r.db.Where(condition).Limit(1).First(&sysLogPo).Error
	if err != nil {
		return sysLogEn, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysLog, err error) {
	sysLogPos := make([]*po.SysLog, 0)
	condition := repositoryimpl.GenerateQueryCondition(queries)

	err = r.db.Find(&sysLogPos, condition).Error
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

	condition = repositoryimpl.GenerateQueryCondition(queries)

	dbQuery := r.db.Model(&po.SysLog{}).Where(condition)

	err := dbQuery.Count(&total).Error
	if err != nil {
		return entries, &rspPag, err
	}

	if total != 0 {
		err = dbQuery.
			Select("sys_log.*").
			Order(reqSort.Sort + " " + reqSort.Direction).
			Scopes(repositoryimpl.Paginate(reqPage.PageNum, reqPage.PageSize)).
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

func (r *SysLog) ExecSql(sql string) error {
	return r.db.Exec(sql).Error
}

func (r *SysLog) FindOneExecSql(sql string) (sysLogEn *entity.SysLog, err error) {
	var sysLogPo po.SysLog
	err = r.db.Raw(sql).Scan(&sysLogPo).Error
	if err != nil {
		return nil, err
	}
	sysLogEn = converter.P2ESysLog(&sysLogPo)

	return
}

func (r *SysLog) FindManyExecSql(sql string) (entries []*entity.SysLog, err error) {
	sysLogPos := make([]*po.SysLog, 0)
	err = r.db.Raw(sql).Scan(&sysLogPos).Error
	if err != nil {
		return nil, err
	}
	entries = converter.P2ESysLogs(sysLogPos)

	return
}
