package ghandler

import (
	sysSvc "github.com/jettjia/go-ddd-demo/application/service/sys"
	grpcGoodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

type GrpcGoodsServer struct {
	grpcGoodsProto.UnimplementedGoodsServer
	SysMenuSrv *sysSvc.SysMenuService
}
