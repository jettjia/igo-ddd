package ginit

import (
	"google.golang.org/grpc"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/interface/grpc/ghandler"
	grpcGoodsProto "jettjia/go-ddd-demo-multi-system/interface/grpc/proto/goods"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(server *grpc.Server, serverApp *cmd.Server) {
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{
		SysMenuSrv: serverApp.Sys,
	})
}
