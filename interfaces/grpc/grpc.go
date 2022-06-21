package grpc

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"github.com/jett/gin-ddd/interfaces/grpc/registersrv"
)

func init() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口") // ip, port 这里应该是自动生成。后续注册到k8s等
	*IP = "127.0.0.1"
	*Port = 5001

	// 启动grpc
	server := grpc.NewServer()

	// 注册pb服务
	registersrv.RegisterSrv(server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	// 启动grpc服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	fmt.Println("启动服务的IP和Port是：", *IP, *Port)
}
