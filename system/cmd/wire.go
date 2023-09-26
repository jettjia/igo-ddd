//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"

	"jettjia/go-ddd-demo-multi-system/application/assembler"
	"jettjia/go-ddd-demo-multi-system/application/service"
	"jettjia/go-ddd-demo-multi-system/config"
	"jettjia/go-ddd-demo-multi-system/domain/aggregate"
	"jettjia/go-ddd-demo-multi-system/domain/srv"
	"jettjia/go-ddd-demo-multi-system/infra/repository/repo"
)

//go:generate wire
func InitServer() (*Server, error) {
	panic(wire.Build(
		config.CfgProvider,

		repo.RepoProviderSet,
		aggregate.AggProviderSet,
		srv.SvcProviderSet,
		assembler.AsseProviderSet,
		service.ServiceProviderSet,

		NewServer,
	),
	)
}
