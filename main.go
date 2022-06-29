package main

import (
	_ "github.com/jettjia/gin-ddd/boot"
	"github.com/jettjia/gin-ddd/cmd"
	"github.com/jettjia/gin-ddd/interfaces/event"
	"github.com/jettjia/gin-ddd/interfaces/grpc"
	"github.com/jettjia/gin-ddd/interfaces/http"
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
