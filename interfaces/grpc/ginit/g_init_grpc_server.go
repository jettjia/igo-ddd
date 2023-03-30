package ginit

import (
	"google.golang.org/grpc"

	"github.com/jettjia/go-ddd-demo/cmd"
	"github.com/jettjia/go-ddd-demo/interfaces/grpc/ghandler"
	grpcGoodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(app *cmd.App, server *grpc.Server) {
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{
		SysMenuSrv: app.SysMenuSvc,
	})
}
