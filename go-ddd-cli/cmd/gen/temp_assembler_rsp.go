package gen

const TempAssemblerTdo = `
package user

import (
	"github.com/jinzhu/copier"

	dtoUser "xdata/xtext/zsvc-example/application/dto/user"
	entityUser "xdata/xtext/zsvc-example/domain/entity/user"
)

// SysMenuDto 请求参数
type SysMenuDto struct {
}

// NewSysMenuDto NewUSysMenuReq
func NewSysMenuDto() *SysMenuDto {
	return &SysMenuDto{}
}

//////////////////////////////////////////////////////////////////
// dto to entity

// D2ECreateSysMenu dto转换成entity
func (a *SysMenuDto) D2ECreateSysMenu(dto *dtoUser.CreateSysMenuReq) *entityUser.SysMenu {
	var rspEn entityUser.SysMenu

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}

	return &rspEn
}

// D2EDeleteSysMenu dto转换成entity
func (a *SysMenuDto) D2EDeleteSysMenu(dto *dtoUser.DelSysMenusReq) *entityUser.SysMenu {
	var rspEn entityUser.SysMenu

	rspEn.Ulid = dto.Ulid

	return &rspEn
}

// D2EUpdateSysMenu dto转换成entity
func (a *SysMenuDto) D2EUpdateSysMenu(dto *dtoUser.UpdateSysMenuReq) *entityUser.SysMenu {
	var rspEn entityUser.SysMenu

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}
	return &rspEn
}

// ////////////////////////////////////////////////////////////////
// entity to dto

// E2DCreateSysMenu dto转换成entity
func (a *SysMenuDto) E2DCreateSysMenu(en *entityUser.SysMenu) (dto *dtoUser.CreateSysMenuRsp) {
	if en == nil {
		return &dtoUser.CreateSysMenuRsp{}
	}
	dto.Ulid = en.Ulid

	return
}

// E2DFindSysMenuRsp entity转换成dto
func (a *SysMenuDto) E2DFindSysMenuRsp(en *entityUser.SysMenu) *dtoUser.FindSysMenuRsp {
	var rspDto dtoUser.FindSysMenuRsp

	if err := copier.Copy(&rspDto, &en); err != nil {
		panic(any(err))
	}

	return &rspDto
}

// E2DGetSysMenus entity转换成dto
func (a *SysMenuDto) E2DGetSysMenus(ens []*entityUser.SysMenu) []*dtoUser.FindSysMenuRsp {
	if len(ens) == 0 {
		return []*dtoUser.FindSysMenuRsp{}
	}

	var SysMenusRsp []*dtoUser.FindSysMenuRsp
	for _, v := range ens {
		cfg := a.E2DFindSysMenuRsp(v)
		SysMenusRsp = append(SysMenusRsp, cfg)
	}

	return SysMenusRsp
}


`
