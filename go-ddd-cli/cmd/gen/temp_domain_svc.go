package gen

const TempDomainSvc = `
package user

import (
	"context"

	"github.com/jettjia/go-pkg/pkg/xsql/builder"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
	repoUser "xdata/xtext/zsvc-example/infra/repository/repo/user"
)

// SysMenuSvc sys_menu_svc
//
//go:generate mockgen --source ./sys_menu_svc.go --destination ./mock/mock_sys_menu_svc.go --package mock
type SysMenuSvc interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error)                                                                                                  // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                               // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                               // 修改
	FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entityUser.SysMenu, err error)                                                                                                // 查看byId
	FindSysMenuByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error)                                                                                // 查看byQuery
	FindSysMenuAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysMenu, err error)                                                                                    // 所有
	DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error)                                                                                                            // delete by
	DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error)                                                                                     // delete by
	FindSysMenuPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysMenu, pageData *builder.PageData, err error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                              // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error)                                                                                                  // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error)                                                                                                 // 原生sql查询多个
}

type SysMenu struct {
	sysMenuRepo *repoUser.SysMenu
}

func NewSysMenuSvc() *SysMenu {
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

func (a *SysMenu) FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindById(ctx, ulid)
}

func (a *SysMenu) FindSysMenuByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindByQuery(ctx, queries)
}

func (a *SysMenu) FindSysMenuAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindAll(ctx, queries)
}

func (a *SysMenu) DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error) {
	return a.sysMenuRepo.DeleteUnscopedByQuery(ctx, queries)
}

func (a *SysMenu) DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error) {
	return a.sysMenuRepo.DeleteByQuery(ctx, queries, sysMenuEn)
}

func (a *SysMenu) FindSysMenuPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysMenu, pageData *builder.PageData, err error) {
	return a.sysMenuRepo.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *SysMenu) ExecSql(ctx context.Context, sql string) error {
	return a.sysMenuRepo.ExecSql(ctx, sql)
}

func (a *SysMenu) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindOneExecSql(ctx, sql)
}

func (a *SysMenu) FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindManyExecSql(ctx, sql)
}

`

const TempDomainSvc2 = `
package user

import (
	"context"

	"github.com/jettjia/go-pkg/pkg/xsql/builder"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
	repoUser "xdata/xtext/zsvc-example/infra/repository/repo/user"
)

// SysMenuSvc sys_menu_svc
//
//go:generate mockgen --source ./sys_menu_svc.go --destination ./mock/mock_sys_menu_svc.go --package mock
type SysMenuSvc interface {
	CreateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error)                                                                                                  // 创建
	DeleteSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                               // 删除
	UpdateSysMenu(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                               // 修改
	FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entityUser.SysMenu, err error)                                                                                                // 查看byId
	FindSysMenuByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error)                                                                                // 查看byQuery
	FindSysMenuAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysMenu, err error)                                                                                    // 所有
	DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error)                                                                                                            // delete by
	DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error)                                                                                    // delete by
	FindSysMenuPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysMenu, pageData *builder.PageData, err error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                              // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error)                                                                                                  // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error)                                                                                                 // 原生sql查询多个
}

type SysMenu struct {
	sysMenuRepo *repoUser.SysMenu
}

func NewSysMenuSvc() *SysMenu {
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

func (a *SysMenu) FindSysMenuById(ctx context.Context, ulid string) (sysMenuEn *entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindById(ctx, ulid)
}

func (a *SysMenu) FindSysMenuByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error) {
	queries = append(queries, &builder.Query{Key: "deleted_at", Value: 0, Operator: builder.Operator_opEq})
	return a.sysMenuRepo.FindByQuery(ctx, queries)
}

func (a *SysMenu) FindSysMenuAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysMenu, err error) {
	queries = append(queries, &builder.Query{Key: "deleted_at", Value: 0, Operator: builder.Operator_opEq})
	return a.sysMenuRepo.FindAll(ctx, queries)
}

func (a *SysMenu) DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error) {
	return a.sysMenuRepo.DeleteUnscopedByQuery(ctx, queries)
}

func (a *SysMenu) DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error) {
	return a.sysMenuRepo.DeleteByQuery(ctx, queries, sysMenuEn)
}

func (a *SysMenu) FindSysMenuPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysMenu, pageData *builder.PageData, err error) {
	queries = append(queries, &builder.Query{Key: "deleted_at", Value: 0, Operator: builder.Operator_opEq})
	return a.sysMenuRepo.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *SysMenu) ExecSql(ctx context.Context, sql string) error {
	return a.sysMenuRepo.ExecSql(ctx, sql)
}

func (a *SysMenu) FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindOneExecSql(ctx, sql)
}

func (a *SysMenu) FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error) {
	return a.sysMenuRepo.FindManyExecSql(ctx, sql)
}

`
