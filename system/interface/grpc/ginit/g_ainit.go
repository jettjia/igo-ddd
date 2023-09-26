package ginit

import (
	"google.golang.org/grpc"

	"jettjia/go-ddd-demo-multi-system/cmd"
)

func GInit(server *grpc.Server, serverApp *cmd.Server) {
	RegisterGrpcSrv(server, serverApp) // 注册服务
	InitGrpcClient(serverApp)          // 初始化client
}
