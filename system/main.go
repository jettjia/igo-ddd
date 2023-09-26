package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"jettjia/go-ddd-demo-multi-system/cmd"
	"jettjia/go-ddd-demo-multi-system/interface/event"
	"jettjia/go-ddd-demo-multi-system/interface/grpc"
	"jettjia/go-ddd-demo-multi-system/interface/http"
	"jettjia/go-ddd-demo-multi-system/interface/job"
)

var wg sync.WaitGroup

func main() {
	// 系统环境变量
	env := flag.String("env", "debug", "configure environment reading")
	flag.Parse()

	os.Setenv("env", *env)

	// 依赖注入
	server, err := cmd.InitServer()
	if err != nil {
		panic(err)
	}

	wg.Add(2)
	// start http
	go func() {
		defer wg.Done()
		http.InitHttp(server)
	}()

	// start grpc
	go func() {
		defer wg.Done()
		grpc.InitGrpc(server)
	}()

	// start event mq
	event.InitEvent()

	// start InitJob
	job.InitJob()

	wg.Wait()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
