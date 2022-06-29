package assembler

import (
	userDto "github.com/jett/gin-ddd/application/dto/user"
	entity "github.com/jett/gin-ddd/domain/entity/user"
)

type UserRSP struct {
}

func NewUserRSP() *UserRSP {
	return &UserRSP{}
}

// E2DSimpleUserInfo 把用户实体映射到简单的实体 dto中，返回给前端
func (u *UserRSP) E2DSimpleUserInfo(user *entity.User) *userDto.SimpleUserInfo {
	simpleUser := &userDto.SimpleUserInfo{}
	simpleUser.Id = user.ID
	simpleUser.Nickname = user.Nickname

	return simpleUser
}
