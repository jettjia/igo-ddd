package main

import (
	_ "github.com/jettjia/go-ddd/boot"
	"github.com/jettjia/go-ddd/cmd"
	"github.com/jettjia/go-ddd/interfaces/event"
	"github.com/jettjia/go-ddd/interfaces/grpc"
	"github.com/jettjia/go-ddd/interfaces/http"
)

func main() {
	app, err := cmd.InitApp()
	if err != nil {
		panic(err)
	}
	http.InitHttp(app)   // start http
	grpc.InitGrpc()      // start grpc
	event.InitEvent(app) // start event mq

	select {}
}
