package ginit

import (
	"google.golang.org/grpc"

	service "github.com/jettjia/go-ddd-demo/application/service/sys"
	"github.com/jettjia/go-ddd-demo/interfaces/grpc/ghandler"
	grpcGoodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(server *grpc.Server) {
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{
		SysMenuSrv: service.NewSysMenuService(),
	})
}
