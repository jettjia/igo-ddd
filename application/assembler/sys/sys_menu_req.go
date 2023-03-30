package assembler

import (
	"github.com/jinzhu/copier"

	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	entity "github.com/jettjia/go-ddd-demo/domain/entity/sys"
	"github.com/jettjia/go-ddd-demo/global"
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
	rspEn.CreatedBy = global.GCustomerInfo.Username

	return &rspEn
}

// D2EDeleteSysMenu dto转换成entity
func (a *SysMenuReq) D2EDeleteSysMenu(dto *dto.DelSysMenusReq) *entity.SysMenu {
	var rspEn entity.SysMenu

	rspEn.Id = dto.Id
	rspEn.DeletedBy = global.GCustomerInfo.Username

	return &rspEn
}

// D2EUpdateSysMenu dto转换成entity
func (a *SysMenuReq) D2EUpdateSysMenu(dto *dto.UpdateSysMenuReq) *entity.SysMenu {
	var rspEn entity.SysMenu

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}
	rspEn.UpdatedBy = global.GCustomerInfo.Username
	return &rspEn
}
