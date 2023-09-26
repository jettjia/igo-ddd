package grpc

import (
	"fmt"

	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/interface/grpc/ginit"
	"jettjia/go-ddd-demo-multi-system/interface/grpc/middleware"
)

func InitGrpc(serverApp *cmd.Server) {
	if !serverApp.Cfg.Server.EnableGrpc {
		return
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middleware.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(middleware.RecoverInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(middleware.RecoverInterceptor()),
		)),
	)

	ginit.GInit(server, serverApp) // 初始化

	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", serverApp.Cfg.Gserver.Host, serverApp.Cfg.Gserver.PublicPort))

	// 启动grpc服务
	go func() {
		err := server.Serve(listener)
		if err != nil {
			panic("InitGrpc:failed to start grpc:" + err.Error())
		}
	}()

	fmt.Printf("[Grpc-debug] Listening and serving RPC on :%d \r\n", serverApp.Cfg.Gserver.PublicPort)
}
