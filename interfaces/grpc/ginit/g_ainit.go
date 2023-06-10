package ginit

import (
	"google.golang.org/grpc"
)

func GInit(server *grpc.Server) {
	RegisterGrpcSrv(server) // 注册服务
	InitGrpcClient()        // 初始化client
}
