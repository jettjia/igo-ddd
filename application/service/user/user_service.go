package service

import (
	"context"

	assembler "github.com/jett/gin-ddd/application/assembler/user"
	dto "github.com/jett/gin-ddd/application/dto/user"
	"github.com/jett/gin-ddd/domain/service/user"
)

type UserService struct {
	assUserRSP    *assembler.UserRSP
	assUserREQ    *assembler.UserREQ
	domainService user.UserService
}

func NewUserService() *UserService {
	return &UserService{
		assUserRSP:    assembler.NewUserRSP(),
		assUserREQ:    assembler.NewUserREQ(),
		domainService: user.NewUserService(),
	}
}

// GetSimpleUserInfo 获取用户信息给到 interfaces
func (u *UserService) GetSimpleUserInfo(ctx context.Context, req *dto.SimpleUserInfoReq) (*dto.SimpleUserInfo, error) {
	userEntity := u.assUserREQ.D2ESimpleUserInfo(req)
	entUser, err := u.domainService.FindUserByID(ctx, userEntity.ID) // 业务复杂的话，这里应该调用 domain/aggregate聚合
	if err != nil {
		return nil, err
	}
	return u.assUserRSP.E2DSimpleUserInfo(entUser), nil
}
