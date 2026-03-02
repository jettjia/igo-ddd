package user

import (
	serviceUser "github.com/jettjia/ddddemo/application/service/user"
)

type Handler struct {
	SysUserSrv *serviceUser.SysUserService
}

func NewHandler() *Handler {
	return &Handler{
		SysUserSrv: serviceUser.NewSysUserService(),
	}
}
