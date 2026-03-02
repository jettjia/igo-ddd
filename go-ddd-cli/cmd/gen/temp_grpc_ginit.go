package gen

const TempGrpcGinit = `

package ginit

import (
	"google.golang.org/grpc"

	grpcGoodsProto "xdata/xtext/common/proto/goods"

	"xdata/xtext/zsvc-example/application/service"
	"xdata/xtext/zsvc-example/api/grpc/ghandler"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(server *grpc.Server) {
	svc := service.NewSysMenuService()
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{
		SysMenuSrv: svc,
	})
}

`
