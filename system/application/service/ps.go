package service

import "github.com/google/wire"

var ServiceProviderSet = wire.NewSet(NewSysMenuService)
