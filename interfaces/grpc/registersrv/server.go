package registersrv

import (
	"google.golang.org/grpc"

	service "github.com/jettjia/gin-ddd/application/service/user"
	"github.com/jettjia/gin-ddd/interfaces/grpc/proto"
	"github.com/jettjia/gin-ddd/interfaces/grpc/server/user"
)

func RegisterSrv(server *grpc.Server) {
	proto.RegisterUserServer(server, &user.UserServer{
		UserSrv: service.NewUserService(),
	})
}
