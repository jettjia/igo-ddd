package user

import (
	"context"

	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
	repoUser "github.com/jettjia/ddddemo/infra/repository/repo/user"
)

// SysUserSvc sys_user_svc
//
//go:generate mockgen --source ./sys_user_svc.go --destination ./mock/mock_sys_user_svc.go --package mock
type SysUserSvc interface {
	CreateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (ulid string, err error)                                                                                                  // 创建
	DeleteSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error)                                                                                                               // 删除
	UpdateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error)                                                                                                               // 修改
	FindSysUserById(ctx context.Context, ulid string) (sysUserEn *entityUser.SysUser, err error)                                                                                                // 查看byId
	FindSysUserByQuery(ctx context.Context, queries []*builder.Query) (sysUserEn *entityUser.SysUser, err error)                                                                                // 查看byQuery
	FindSysUserAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysUser, err error)                                                                                    // 所有
	FindSysUserPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysUser, pageData *builder.PageData, err error) // 列表
	ExecSql(ctx context.Context, sql string) error                                                                                                                                              // 执行sql
	FindOneExecSql(ctx context.Context, sql string) (sysUserEn *entityUser.SysUser, err error)                                                                                                  // 原生sql查询单个
	FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysUser, err error)                                                                                                 // 原生sql查询多个
}

type SysUser struct {
	sysUserRepo *repoUser.SysUser
}

func NewSysUserSvc() *SysUser {
	return &SysUser{
		sysUserRepo: repoUser.NewSysUserImpl(),
	}
}

func (a *SysUser) CreateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (ulid string, err error) {
	return a.sysUserRepo.Create(ctx, sysUserEn)
}

func (a *SysUser) DeleteSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error) {
	return a.sysUserRepo.Delete(ctx, sysUserEn)
}

func (a *SysUser) UpdateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error) {
	return a.sysUserRepo.Update(ctx, sysUserEn)
}

func (a *SysUser) FindSysUserById(ctx context.Context, ulid string) (sysUserEn *entityUser.SysUser, err error) {
	return a.sysUserRepo.FindById(ctx, ulid)
}

func (a *SysUser) FindSysUserByQuery(ctx context.Context, queries []*builder.Query) (sysUserEn *entityUser.SysUser, err error) {
	return a.sysUserRepo.FindByQuery(ctx, queries)
}

func (a *SysUser) FindSysUserAll(ctx context.Context, queries []*builder.Query) (entries []*entityUser.SysUser, err error) {
	return a.sysUserRepo.FindAll(ctx, queries)
}

func (a *SysUser) FindSysUserPage(ctx context.Context, queries []*builder.Query, reqPage *builder.PageData, reqSort *builder.SortData) (entries []*entityUser.SysUser, pageData *builder.PageData, err error) {
	return a.sysUserRepo.FindPage(ctx, queries, reqPage, reqSort)
}

func (a *SysUser) ExecSql(ctx context.Context, sql string) error {
	return a.sysUserRepo.ExecSql(ctx, sql)
}

func (a *SysUser) FindOneExecSql(ctx context.Context, sql string) (sysUserEn *entityUser.SysUser, err error) {
	return a.sysUserRepo.FindOneExecSql(ctx, sql)
}

func (a *SysUser) FindManyExecSql(ctx context.Context, sql string) (entries []*entityUser.SysUser, err error) {
	return a.sysUserRepo.FindManyExecSql(ctx, sql)
}
