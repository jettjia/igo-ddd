package gen

const TempDomainIrepository = `
package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
)

// ISysMenuRepo sys_menu
//
//go:generate mockgen --source ./i_sys_menu_repo.go --destination ./mock/mock_i_sys_menu_repo.go --package mock
type ISysMenuRepo interface {
	Create(ctx context.Context, sysMenuEn *entityUser.SysMenu) (ulid string, err error)                                                                                                     // 创建
	Delete(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                                  // 删除
	Update(ctx context.Context, sysMenuEn *entityUser.SysMenu) (err error)                                                                                                                  // 修改
	FindById(ctx context.Context, ulid string, selectColumn ...string) (sysMenuEn *entityUser.SysMenu, err error)                                                                           // 查看byId
	FindByQuery(ctx context.Context, queries []*builder.Query) (sysMenuEn *entityUser.SysMenu, err error)                                                                                   // 查看byQuery
	FindAll(ctx context.Context, queries []*builder.Query, selectArgs ...[]string) (entries []*entityUser.SysMenu, err error)                                                               // 所有
	DeleteUnscopedByQuery(ctx context.Context, queries []*builder.Query) (err error)                                                                                                        // delete by
	DeleteByQuery(ctx context.Context, queries []*builder.Query, sysMenuEn *entityUser.SysMenu) (err error)                                                                                 // delete by
	FindPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData, selectArgs ...[]string) ([]*entityUser.SysMenu, *builder.PageData, error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                          // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysMenuEn *entityUser.SysMenu, err error)                                                                                              // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysMenu, err error)                                                                                             // 原生sql查询多个
}
`
