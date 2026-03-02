package subscribe

import (
	serviceUser "github.com/jettjia/ddddemo/application/service/user"
)

type Subscribe struct {
	SysUserSrv *serviceUser.SysUserService
}

func NewSubscribe() *Subscribe {
	return &Subscribe{
		SysUserSrv: serviceUser.NewSysUserService(),
	}
}
