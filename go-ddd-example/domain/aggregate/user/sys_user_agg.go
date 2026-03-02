package user

import (
	"context"
	"fmt"

	"github.com/jettjia/igo-pkg/pkg/data"
	"github.com/jettjia/igo-pkg/pkg/xerror"
	"github.com/jettjia/igo-pkg/pkg/xsql/builder"

	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
	"github.com/jettjia/ddddemo/infra/pkg/idata"
	repoUser "github.com/jettjia/ddddemo/infra/repository/repo/user"
	"github.com/jettjia/ddddemo/types/apierror"
)

// SysUserAgg sys_user_agg
//
//go:generate mockgen --source ./sys_user_agg.go --destination ./mock/mock_sys_user_agg.go --package mock
type SysUserAgg interface {
	CreateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (ulid string, err error)   // 创建
	DeleteSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error)                // 删除
	UpdateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error)                // 修改
	FindSysUserById(ctx context.Context, ulid string) (sysUserEn *entityUser.SysUser, err error) // 查看byId
	Logout(ctx context.Context, token string) (err error)                                        // 登出
}

type SysUser struct {
	tx          data.Transaction  // 事务管理器
	sysUserRepo *repoUser.SysUser // db user
	sysLogRepo  *repoUser.SysLog  // db sys_log
}

func NewSysUserAgg() *SysUser {
	return &SysUser{
		tx:          idata.NewDataOptionCli(),
		sysUserRepo: repoUser.NewSysUserImpl(),
		sysLogRepo:  repoUser.NewSysLogImpl(),
	}
}

func (a *SysUser) CreateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (ulid string, err error) {
	var (
		sysUserEnDB *entityUser.SysUser
	)

	// 调用事务实例
	err = a.tx.ExecTx(ctx, func(ctx context.Context) error {
		// 判断用户昵称是否存在
		queries := []*builder.Query{{Key: "nick_name", Value: sysUserEn.NickName, Operator: builder.Operator_opEq}}
		if sysUserEnDB, err = a.sysUserRepo.FindByQuery(ctx, queries); err != nil {
			return err
		}
		if sysUserEnDB.Ulid != "" {
			err = xerror.NewErrorOpt(
				apierror.UserNameConflictErr,
				xerror.WithCause(fmt.Sprintf("the nick_name: %s has", sysUserEn.NickName)),
				xerror.WithSolution("please change the nick_name"),
			)
			return err
		}
		// db 创建用户
		if ulid, err = a.sysUserRepo.Create(ctx, sysUserEn); err != nil {
			return err
		}

		return err
	})

	return
}

func (a *SysUser) DeleteSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error) {
	err = a.tx.ExecTx(ctx, func(ctx context.Context) error {
		err = a.sysUserRepo.Delete(ctx, sysUserEn)

		var sysLogEn entityUser.SysLog
		sysLogEn.CreatedBy = sysUserEn.DeletedBy
		sysLogEn.Msg = "SysUser.Delete"
		_, err = a.sysLogRepo.Create(ctx, &sysLogEn)

		return err
	})

	return
}

func (a *SysUser) UpdateSysUser(ctx context.Context, sysUserEn *entityUser.SysUser) (err error) {
	err = a.tx.ExecTx(ctx, func(ctx context.Context) error {
		err = a.sysUserRepo.Update(ctx, sysUserEn)

		var sysLogEn entityUser.SysLog
		sysLogEn.CreatedBy = sysUserEn.DeletedBy
		sysLogEn.Msg = "SysUser.Delete"
		_, err = a.sysLogRepo.Create(ctx, &sysLogEn)

		return err
	})

	return
}

// FindSysUserById 获取用户信息，同时带上用户级别信息
func (a *SysUser) FindSysUserById(ctx context.Context, ulid string) (sysUserEn *entityUser.SysUser, err error) {
	if sysUserEn, err = a.sysUserRepo.FindById(ctx, ulid); err != nil {
		return
	}

	// 这里定义的错误返回是，用户未找到，由自己定义的
	if sysUserEn.Ulid == "" {
		err = xerror.NewError(apierror.UserNotFoundErr, fmt.Sprintf("user_id: %s, info not found", ulid), nil)
		return
	}

	return
}
