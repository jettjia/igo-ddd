package ghandler

import (
	grpcGoodsProto "github.com/jettjia/ddddemo/idl/proto/goods"
)

type GrpcGoodsServer struct {
	grpcGoodsProto.UnimplementedGoodsServer
}
