package ghandler

import (
	sysSvc "jettjia/go-ddd-demo-multi-system/application/service"
	grpcGoodsProto "jettjia/go-ddd-demo-multi-system/interface/grpc/proto/goods"
)

type GrpcGoodsServer struct {
	grpcGoodsProto.UnimplementedGoodsServer
	SysMenuSrv *sysSvc.SysMenuService
}
