package service

import (
	"fmt"
	repository "github.com/jett/gin-ddd/domain/repository/user"
)

type UserLoginService struct {
	userRepo repository.IUserRepository
}

func (this *UserLoginService) Login(userName string, userPwd string) (string, error) {
	user, err := this.userRepo.GetUserByName(userName)
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
