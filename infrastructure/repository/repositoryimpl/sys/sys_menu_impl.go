package repositoryimpl

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	isysMenu "github.com/jettjia/go-ddd-demo/domain/irepository/sys"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
	converter "github.com/jettjia/go-ddd-demo/infrastructure/repository/converter/sys"
	po "github.com/jettjia/go-ddd-demo/infrastructure/repository/po/sys"
	"github.com/jettjia/go-ddd-demo/infrastructure/repository/repositoryimpl"
	"github.com/jettjia/go-ddd-demo/types"
)

var _ isysMenu.ISysMenuRepo = (*SysMenu)(nil)

type SysMenu struct {
	db *gorm.DB
}

func NewSysMenuImpl() *SysMenu {
	return &SysMenu{
		db: global.GDB,
	}
}

func (r *SysMenu) Create(ctx context.Context, sysMenuEn *entity.SysMenu) (id uint64, err error) {
	sysMenuPo := converter.E2PSysMenuAdd(sysMenuEn)
	err = r.db.Create(&sysMenuPo).Error
	if err != nil {
		return 0, err
	}

	return sysMenuPo.Id, nil
}

func (r SysMenu) Delete(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	sysMenuPo := converter.E2PSysMenuDel(sysMenuEn)

	return r.db.Model(&po.SysMenu{}).Where("id = ? ", sysMenuEn.Id).Updates(sysMenuPo).Error
}

func (r *SysMenu) Update(ctx context.Context, sysMenuEn *entity.SysMenu) (err error) {
	sysMenuPo := converter.E2PSysMenuUpdate(sysMenuEn)

	return r.db.Model(&po.SysMenu{}).Where("id = ? ", sysMenuEn.Id).Updates(sysMenuPo).Error
}

func (r *SysMenu) FindById(ctx context.Context, id uint64) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	err = r.db.Limit(1).First(&sysMenuPo, "id = ? ", id).Error
	if err != nil {
		return sysMenuEn, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindByQuery(ctx context.Context, queries []*types.Query) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	condition := repositoryimpl.GenerateQueryCondition(queries)
	err = r.db.Where(condition).Limit(1).First(&sysMenuPo).Error
	if err != nil {
		return sysMenuEn, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindAll(ctx context.Context, queries []*types.Query) (entries []*entity.SysMenu, err error) {
	sysMenuPos := make([]*po.SysMenu, 0)
	condition := repositoryimpl.GenerateQueryCondition(queries)

	err = r.db.Find(&sysMenuPos, condition).Error
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

	condition = repositoryimpl.GenerateQueryCondition(queries)

	dbQuery := r.db.Model(&po.SysMenu{}).Where(condition)

	err := dbQuery.Count(&total).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		err = gerror.NewCode(responseutil.CommInternalServer, err.Error())

		return entries, &rspPag, err
	}

	if total != 0 {
		err := dbQuery.
			Select("sys_menu.*").
			Order(reqSort.Sort + " " + reqSort.Direction).
			Scopes(repositoryimpl.Paginate(reqPage.PageNum, reqPage.PageSize)).
			Find(&sysMenuPos).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
			err = gerror.NewCode(responseutil.CommInternalServer, err.Error())

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

func (r *SysMenu) ExecSql(sql string) error {
	return r.db.Exec(sql).Error
}

func (r *SysMenu) FindOneExecSql(sql string) (sysMenuEn *entity.SysMenu, err error) {
	var sysMenuPo po.SysMenu
	err = r.db.Raw(sql).Scan(&sysMenuPo).Error
	if err != nil {
		return nil, err
	}
	sysMenuEn = converter.P2ESysMenu(&sysMenuPo)

	return
}

func (r *SysMenu) FindManyExecSql(sql string) (entries []*entity.SysMenu, err error) {
	sysMenuPos := make([]*po.SysMenu, 0)
	err = r.db.Raw(sql).Scan(&sysMenuPos).Error
	if err != nil {
		return nil, err
	}
	entries = converter.P2ESysMenus(sysMenuPos)

	return
}
