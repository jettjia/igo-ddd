package grpc

import (
	"fmt"
	"github.com/jettjia/go-ddd-demo/boot"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jettjia/go-ddd-demo/interfaces/grpc/middleware"
	"google.golang.org/grpc"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/interfaces/grpc/ginit"
)

func InitGrpc(app *boot.App) {
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

	ginit.GInit(app, server) // 初始化

	listener, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", global.Gconfig.Gserver.Host, global.Gconfig.Gserver.PublicPort))

	// 启动grpc服务
	go func() {
		err := server.Serve(listener)
		if err != nil {
			panic("InitGrpc:failed to start grpc:" + err.Error())
		}
	}()

	fmt.Printf("[Grpc-debug] Listening and serving HTTP on :%d \r\n", global.Gconfig.Gserver.PublicPort)
}
