package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/jett/gin-ddd/interfaces/grpc/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())

	defer conn.Close()

	c := proto.NewUserClient(conn)
	rsp, _ := c.FindSimpleUser(context.Background(),
		&proto.FindSimpleUserRequest{Id: 1},
	)

	fmt.Println(rsp.Nickname)
}
