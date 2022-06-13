//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"

	userService "github.com/jett/gin-ddd/application/service/user"
	"github.com/jett/gin-ddd/interfaces"
)

//go:generate wire
var providerSet = wire.NewSet(
	// 注入 application
	userService.NewUserService,

	// 注入项目启动入口
	interfaces.NewServer,
)

func NewApp() (*interfaces.Server, error) {
	panic(wire.Build(providerSet))
}
