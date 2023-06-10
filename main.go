package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/go-ddd-demo/boot"
	"github.com/jettjia/go-ddd-demo/interfaces/event"
	"github.com/jettjia/go-ddd-demo/interfaces/grpc"
	"github.com/jettjia/go-ddd-demo/interfaces/http"
	"github.com/jettjia/go-ddd-demo/interfaces/job"
)

func main() {
	ENV := flag.String("env", "debug", "环境,配置读取")
	flag.Parse()

	// 全局配置
	boot.InitServer(*ENV)

	// start http
	http.InitHttp()

	// start grpc
	grpc.InitGrpc()

	// start event mq
	event.InitEvent()

	// start InitJob
	job.InitJob(*ENV)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
