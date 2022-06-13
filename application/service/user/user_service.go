package service

import (
	assembler "github.com/jett/gin-ddd/application/assembler/user"
	dto "github.com/jett/gin-ddd/application/dto/user"
	"github.com/jett/gin-ddd/domain/service"
)

type UserService struct {
	assUserRSP    *assembler.UserRSP
	assUserREQ    *assembler.UserREQ
	domainService service.UserService
}

func NewUserService() *UserService {
	return &UserService{
		assUserRSP:    assembler.NewUserRSP(),
		assUserREQ:    assembler.NewUserREQ(),
		domainService: *service.NewUserService(),
	}
}

// GetSimpleUserInfo 获取用户信息给到 interfaces
func (this *UserService) GetSimpleUserInfo(req *dto.SimpleUserInfoReq) *dto.SimpleUserInfo {
	userEntity := this.assUserREQ.D2E_SimpleUserInfo(req)
	entUser, _ := this.domainService.GetUser(userEntity.ID) // 业务复杂的话，这里应该调用 domain/aggregate聚合

	return this.assUserRSP.E2D_SimpleUserInfo(entUser)
}
