package ginit

import (
	"fmt"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/consts"
	grpcGoodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
	"google.golang.org/grpc"
)

var (
	// GoodsClient grpc客户端,goods服务
	GoodsClient grpcGoodsProto.GoodsClient
)

// InitGrpcClient 初始化链接其他服务的client
func InitGrpcClient() {
	conf := global.Gconfig.Gserver
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.ClientGoodsHost, conf.ClientGoodsPort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(conf.MaxMsgSize*consts.UnitM)),
	)
	if err != nil {
		panic("InitGrpcClient:err:" + err.Error())
	}

	GoodsClient = grpcGoodsProto.NewGoodsClient(conn)
}
