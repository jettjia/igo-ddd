package main

import (
	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/cmd"
	"github.com/jett/gin-ddd/interfaces/event"
	"github.com/jett/gin-ddd/interfaces/grpc"
	"github.com/jett/gin-ddd/interfaces/http"
)

func main() {
	app, err := cmd.InitApp()
	if err != nil {
		panic(err)
	}
	http.InitHttp(app)   // start http
	grpc.InitGrpc()      // start grpc
	event.InitEvent(app) // start event mq
}
