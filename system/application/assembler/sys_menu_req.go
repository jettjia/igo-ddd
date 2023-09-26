package assembler

import (
	"github.com/jinzhu/copier"

	"jettjia/go-ddd-demo-multi-system/application/dto"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

// SysMenuReq 请求参数
type SysMenuReq struct {
}

// NewSysMenuReq NewUSysMenuReq
func NewSysMenuReq() *SysMenuReq {
	return &SysMenuReq{}
}

// D2ECreateSysMenu dto转换成entity
func (a *SysMenuReq) D2ECreateSysMenu(dto *dto.CreateSysMenuReq) *entity.SysMenu {
	var rspEn entity.SysMenu

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}

	return &rspEn
}

// D2EDeleteSysMenu dto转换成entity
func (a *SysMenuReq) D2EDeleteSysMenu(dto *dto.DelSysMenusReq) *entity.SysMenu {
	var rspEn entity.SysMenu

	rspEn.Ulid = dto.Ulid

	return &rspEn
}

// D2EUpdateSysMenu dto转换成entity
func (a *SysMenuReq) D2EUpdateSysMenu(dto *dto.UpdateSysMenuReq) *entity.SysMenu {
	var rspEn entity.SysMenu

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}
	return &rspEn
}
