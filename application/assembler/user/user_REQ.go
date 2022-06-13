package assembler

import (
	"github.com/go-playground/validator/v10"
	entity "github.com/jett/gin-ddd/domain/entity/user"

	dto "github.com/jett/gin-ddd/application/dto/user"
)

type UserREQ struct {
	v *validator.Validate
}

// D2E_SimpleUserInfo 将传入的 dto的SimpleUserInfoReq转成 entity.User
func (this *UserREQ) D2E_SimpleUserInfo(dto *dto.SimpleUserInfoReq) *entity.User {
	var userEntity entity.User
	err := this.v.Struct(dto)
	if err != nil {
		panic(err) // 临时处理
	}

	userEntity.ID = dto.Id

	return &userEntity
}
