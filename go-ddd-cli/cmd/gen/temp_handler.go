package gen

const TempHandler = `
package user

import (
	serviceUser "xdata/xtext/zsvc-example/application/service/user"
)

type Handler struct {
	SysMenuSrv *serviceUser.SysMenuService
}

func NewHandler() *Handler {
	return &Handler{
		SysMenuSrv: serviceUser.NewSysMenuService(),
	}
}
`
