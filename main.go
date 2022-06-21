package main

import (
	_ "github.com/jett/gin-ddd/boot"
	_ "github.com/jett/gin-ddd/interfaces/grpc"
	_ "github.com/jett/gin-ddd/interfaces/http"
)

func main() {
	select {}
}
