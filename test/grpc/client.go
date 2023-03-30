package grpc

import (
	"google.golang.org/grpc"

	goodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

var (
	GrpcClient goodsProto.GoodsClient
	ClientConn *grpc.ClientConn
)

func Init() {
	var err error
	ClientConn, err = grpc.Dial("0.0.0.0:18080", grpc.WithInsecure())
	if err != nil {
		panic(any(err))
	}
	GrpcClient = goodsProto.NewGoodsClient(ClientConn)
}
