package ginit

import (
	"google.golang.org/grpc"

	grpcGoodsProto "github.com/jettjia/ddddemo/idl/proto/goods"

	"github.com/jettjia/ddddemo/api/grpc/ghandler"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(server *grpc.Server) {
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{})
}
