package ginit

import (
	"fmt"

	"google.golang.org/grpc"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/infra/consts"
	grpcGoodsProto "jettjia/go-ddd-demo-multi-system/interface/grpc/proto/goods"
)

var (
	// GoodsClient grpc客户端,goods服务
	GoodsClient grpcGoodsProto.GoodsClient
)

// InitGrpcClient 初始化链接其他服务的client
func InitGrpcClient(serverApp *cmd.Server) {
	conf := serverApp.Cfg.Gserver
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", conf.ClientGoodsHost, conf.ClientGoodsPort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(conf.MaxMsgSize*consts.UnitM)),
	)
	if err != nil {
		panic("InitGrpcClient:err:" + err.Error())
	}

	GoodsClient = grpcGoodsProto.NewGoodsClient(conn)
}
