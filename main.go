package main

import (
	"flag"

	"github.com/jettjia/go-ddd-demo/boot"
	"github.com/jettjia/go-ddd-demo/cmd"
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

	// 依赖注入
	app := cmd.InitApp()

	// start http
	http.InitHttp(app)

	// start grpc
	grpc.InitGrpc(app)

	// start event mq
	event.InitEvent(app)

	// start InitJob
	job.InitJob(app, *ENV)

	select {}
}
