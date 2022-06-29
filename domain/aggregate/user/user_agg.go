package aggregate

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	repository "github.com/jett/gin-ddd/domain/irepository/user"
)

type UserAgg interface {
	AddUser(context.Context) error // 添加用户
}

var _ UserAgg = (*userAgg)(nil)

// UserAgg 会有聚合：用户+用户日志
type userAgg struct {
	user        *entity.User
	userLog     *entity.UserLog
	userRepo    repository.IUserRepository
	userRepoLog repository.IUserLogRepository
}

// AddUser 创建会员,创建的时候会增加日志
func (u *userAgg) AddUser(ctx context.Context) error {
	entityUser, err := u.userRepo.SaveUser(ctx, u.user)
	if err != nil {
		return err

	}

	u.userLog.UserId = entityUser.ID
	u.userLog.Log = "创建用户" + entityUser.Nickname
	_, err = u.userRepoLog.SaveLog(ctx, u.userLog)
	if err != nil {
		return err
	}
	return nil
}

// Create 创建会员,创建的时候会增加日志
func (u *userAgg) Create(ctx context.Context) error {
	entityUser, err := u.userRepo.SaveUser(ctx, u.user)
	if err != nil {
		return err

	}

	u.userLog.UserId = entityUser.ID
	u.userLog.Log = "创建用户:" + entityUser.Nickname
	_, err = u.userRepoLog.SaveLog(ctx, u.userLog)
	if err != nil {
		return err
	}
	return nil
}
