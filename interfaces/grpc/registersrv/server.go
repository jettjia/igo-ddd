package registersrv

import (
	"google.golang.org/grpc"

	service "github.com/jett/gin-ddd/application/service/user"
	"github.com/jett/gin-ddd/interfaces/grpc/proto"
	"github.com/jett/gin-ddd/interfaces/grpc/server/user"
)

func RegisterSrv(server *grpc.Server) {
	proto.RegisterUserServer(server, &user.UserServer{
		UserSrv: service.NewUserService(),
	})
}
