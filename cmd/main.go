package main

import (
	"github.com/jett/gin-ddd/interfaces/http"

	_ "github.com/jett/gin-ddd/boot"
	"github.com/jett/gin-ddd/interfaces/http/handler/user"
)

func main() {
	server := &http.NewHttpApp{
		URESTHandler: user.NewRESTHandler(),
	}
	server.Start()

	select {}
}
