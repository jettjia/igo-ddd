package ginit

import (
	"fmt"

	"google.golang.org/grpc"

	grpcGoodsProto "github.com/jettjia/ddddemo/idl/proto/goods"

	"github.com/jettjia/ddddemo/config"
	"github.com/jettjia/ddddemo/types/consts"
)

var (
	// GoodsClient grpc客户端,goods服务
	GoodsClient grpcGoodsProto.GoodsClient
)

// InitGrpcClient 初始化链接其他服务的client
func InitGrpcClient() {

	cfg := config.NewConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Gserver.ClientGoodsHost, cfg.Gserver.ClientGoodsPort),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(cfg.Gserver.MaxMsgSize*consts.UnitM)),
	)
	if err != nil {
		panic("InitGrpcClient:err:" + err.Error())
	}

	GoodsClient = grpcGoodsProto.NewGoodsClient(conn)
}
