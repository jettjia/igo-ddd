package service

import (
	assembler "github.com/jett/gin-ddd/application/assembler/user"
	dto "github.com/jett/gin-ddd/application/dto/user"
	repository "github.com/jett/gin-ddd/domain/repository/user"
)

type UserService struct {
	assUserRSP *assembler.UserRSP
	assUserREQ *assembler.UserREQ
	userRepo   repository.IUserRepository
}

func NewUserService(assUserRSP *assembler.UserRSP, assUserREQ *assembler.UserREQ, userRepo repository.IUserRepository) *UserService {
	return &UserService{assUserRSP: assUserRSP, assUserREQ: assUserREQ, userRepo: userRepo}
}

// GetSimpleUserInfo 获取用户信息给到 interfaces
func (this *UserService) GetSimpleUserInfo(req *dto.SimpleUserInfoReq) *dto.SimpleUserInfo {
	userEntity := this.assUserREQ.D2E_SimpleUserInfo(req)
	entUser, _ := this.userRepo.GetUser(userEntity.ID) // 业务复杂的话，这里应该调用 domain/aggregate聚合

	return this.assUserRSP.E2D_SimpleUserInfo(entUser)
}
