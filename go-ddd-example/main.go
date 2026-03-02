package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/jettjia/ddddemo/api/event"
	"github.com/jettjia/ddddemo/api/grpc"
	"github.com/jettjia/ddddemo/api/http"
	"github.com/jettjia/ddddemo/api/job"
	"github.com/jettjia/ddddemo/boot"
	"github.com/jettjia/ddddemo/infra/repository/po"
)

func main() {
	// create shutdown channel
	shutdown := make(chan struct{})

	// init config
	env := flag.String("env", "debug", "configure environment reading")
	flag.Parse()

	err := os.Setenv("env", *env)
	if err != nil {
		panic(err)
	}

	// auto create table
	if err = po.AutoTable(); err != nil {
		panic(err)
	}

	// init data
	if err := boot.InitData(); err != nil {
		panic(err)
	}

	// start http
	http.InitHttp()

	// start grpc
	grpc.InitGrpc()

	// start event mq
	event.InitEvent()

	// start InitJob
	go func() {
		job.InitJob(shutdown)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// close
	close(shutdown)
}
