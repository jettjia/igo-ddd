package public

import (
	"jettjia/go-ddd-demo-multi-system/application/service"
)

type Handler struct {
	SysMenuSrv *service.SysMenuService
}

func NewHandler(sysMenuSrv *service.SysMenuService) *Handler {
	return &Handler{
		SysMenuSrv: sysMenuSrv,
	}
}
