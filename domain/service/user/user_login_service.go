package user

import (
	"context"
	"fmt"

	repository "github.com/jett/gin-ddd/domain/repository/user"
	dao "github.com/jett/gin-ddd/infrastructure/dao/user"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: dao.NewUserRepo(),
	}
}

func (u *UserService) Login(ctx context.Context, userName string, userPwd string) (string, error) {
	user, err := u.userRepo.GetUserByName(ctx, userName)
	if err != nil {
		return "", err
	}

	if user.ID > 0 { // 有这个用户
		// 判断密码，等等多个操作
		return "1000200", nil
	} else {
		return "1000404", fmt.Errorf("用户不存在")
	}

}
