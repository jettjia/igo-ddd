package grpc

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc/metadata"

	goodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

func Test_CreateSysMenu(t *testing.T) {
	Init()

	c := context.Background()

	md := metadata.MD{}
	md.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzAzODU3MTMsImlzcyI6Ikdhc1BjIiwiVG9rZW5UeXBlIjoibGV2ZWwxIiwidXNlcl9pZCI6MSwicGhvbmUiOiIxNTAyNjYzNTg1OCIsInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlX2lkIjoxfQ.DwPK5so9g6-g6aTFiC13HvolKZMmK78ylnR_GHwx7gg")
	ctx := metadata.NewOutgoingContext(c, md)

	rsp, err := GrpcClient.CreateSysMenu(ctx, &goodsProto.CreateSysMenuReq{
		MenuName:    "MenuName2",
		Desc:        "Desc",
		Route:       "Route",
		State:       1,
		Pid:         0,
		Pname:       "Pname",
		SortOrder:   50,
		BackendType: 1,
	})

	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(rsp)

	defer ClientConn.Close()
}
