package cmd

import (
	"jettjia/go-ddd-demo-multi-common/pkg/conf"
	"jettjia/go-ddd-demo-multi-system/application/service"
)

type Server struct {
	Cfg *conf.Config
	Sys *service.SysMenuService
}

func NewServer(sys *service.SysMenuService, cfg *conf.Config) *Server {
	return &Server{
		Cfg: cfg,
		Sys: sys,
	}
}
