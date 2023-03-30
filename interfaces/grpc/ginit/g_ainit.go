package ginit

import (
	"github.com/jettjia/go-ddd-demo/cmd"
	"google.golang.org/grpc"
)

func GInit(app *cmd.App, server *grpc.Server) {
	RegisterGrpcSrv(app, server) // 注册服务
	InitGrpcClient()             // 初始化client
}
