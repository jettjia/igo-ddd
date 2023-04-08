package ginit

import (
	"github.com/jettjia/go-ddd-demo/boot"
	"google.golang.org/grpc"
)

func GInit(app *boot.App, server *grpc.Server) {
	RegisterGrpcSrv(app, server) // 注册服务
	InitGrpcClient()             // 初始化client
}
