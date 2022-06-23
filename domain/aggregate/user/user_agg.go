package aggregate

import (
	"context"

	entity "github.com/jett/gin-ddd/domain/entity/user"
	repository "github.com/jett/gin-ddd/domain/repository/user"
)

// User 会有聚合：用户+用户日志
type User struct {
	User        *entity.User
	UserLog     *entity.UserLog
	userRepo    repository.IUserRepository
	userRepoLog repository.IUserLogRepository
}

// Create 创建会员,创建的时候会增加日志
func (this *User) Create(ctx context.Context) error {
	entityUser, err := this.userRepo.SaveUser(ctx, this.User)
	if err != nil {
		return err

	}

	this.UserLog.UserId = entityUser.ID
	this.UserLog.Log = "创建用户" + entityUser.Nickname
	_, err = this.userRepoLog.SaveLog(ctx, this.UserLog)
	if err != nil {
		return err
	}
	return nil
}

func (this *User) GetLogs() (ret []*entity.UserLog) {
	return
}
