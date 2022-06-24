package user

import (
	"context"
	dto "github.com/jett/gin-ddd/application/dto/user"
	"github.com/jett/gin-ddd/interfaces/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *UserServer) Find_SimpleUser(ctx context.Context, request *proto.FindSimpleUserRequest) (*proto.FindSimpleUserReply, error) {
	//fmt.Println("i am in......")
	simpleUserReq := &dto.SimpleUserInfoReq{}
	simpleUserReq.Id = request.Id

	//TODO implement me， 这里去调用  application/service
	dtoSimpleUserInfo, err := u.UserSrv.GetSimpleUserInfo(ctx, simpleUserReq)
	if err != nil {
		return nil, status.Error(codes.Aborted, "user not found")
	}

	rest := proto.FindSimpleUserReply{
		Id:       dtoSimpleUserInfo.Id,
		Nickname: dtoSimpleUserInfo.Nickname,
	}

	return &rest, nil
}
