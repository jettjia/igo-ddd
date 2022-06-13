package assembler

import (
	entity "github.com/jett/gin-ddd/domain/entity/user"

	dto "github.com/jett/gin-ddd/application/dto/user"
)

type UserREQ struct {
}

func NewUserREQ() *UserREQ {
	return &UserREQ{}
}

// D2E_SimpleUserInfo 将传入的 dto的SimpleUserInfoReq转成 entity.User
func (this *UserREQ) D2E_SimpleUserInfo(dto *dto.SimpleUserInfoReq) *entity.User {
	var userEntity entity.User

	userEntity.ID = dto.Id

	return &userEntity
}
