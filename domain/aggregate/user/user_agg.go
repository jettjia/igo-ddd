package aggregate

import (
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

func NewUser(user *entity.User, userRepo repository.IUserRepository, userRepoLog repository.IUserLogRepository) *User {
	return &User{User: user, userRepo: userRepo, userRepoLog: userRepoLog}
}

func NewUserById(id uint64, userRepo repository.IUserRepository, userRepoLog repository.IUserLogRepository) *User {
	user, _ := userRepo.GetUser(id)
	return &User{User: user, userRepo: userRepo, userRepoLog: userRepoLog}
}

// Create 创建会员
func (this *User) Create() error {
	entityUser, err := this.userRepo.SaveUser(this.User)
	if err != nil {
		return err

	}

	this.UserLog.UserId = entityUser.ID
	this.UserLog.Log = "创建用户" + entityUser.Nickname
	_, err = this.userRepoLog.SaveLog(this.UserLog)
	if err != nil {
		return err
	}
	return nil
}

func (this *User) GetLogs() (ret []*entity.UserLog) {
	return
}
