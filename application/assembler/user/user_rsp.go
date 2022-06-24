package assembler

import (
	userDto "github.com/jett/gin-ddd/application/dto/user"
	aggregate "github.com/jett/gin-ddd/domain/aggregate/user"
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

func (u *UserRSP) E2DUserInfo(user *aggregate.UserAgg) *userDto.UserInfo {
	userInfo := &userDto.UserInfo{}
	userInfo.Id = user.User.ID
	userInfo.Nickname = user.User.Nickname
	userInfo.Passport = user.User.Passport
	userInfo.Logs = u.E2DUserLogs(user.GetLogs())

	return userInfo
}

// E2DUserLogs todo 完成日志查询
func (u *UserRSP) E2DUserLogs(logs []*entity.UserLog) (ret []*userDto.UserLog) {

	return
}
