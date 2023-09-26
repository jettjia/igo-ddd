package assembler

import (
	"github.com/jinzhu/copier"

	"jettjia/go-ddd-demo-multi-system/application/dto"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

// SysMenuRsp 请求参数
type SysMenuRsp struct {
}

// NewSysMenuRsp NewUSysMenuRsp
func NewSysMenuRsp() *SysMenuRsp {
	return &SysMenuRsp{}
}

// D2ECreateSysMenu dto转换成entity
func (a *SysMenuRsp) D2ECreateSysMenu(en *entity.SysMenu) (dto *dto.CreateSysMenuRsp) {
	dto.Ulid = en.Ulid

	return
}

// E2DFindSysMenuRsp entity转换成dto
func (a *SysMenuRsp) E2DFindSysMenuRsp(en *entity.SysMenu) *dto.FindSysMenuRsp {
	var rspDto dto.FindSysMenuRsp

	if err := copier.Copy(&rspDto, &en); err != nil {
		panic(any(err))
	}

	return &rspDto
}

// E2DGetSysMenus entity转换成dto
func (a *SysMenuRsp) E2DGetSysMenus(ens []*entity.SysMenu) []*dto.FindSysMenuRsp {
	if len(ens) == 0 {
		return []*dto.FindSysMenuRsp{}
	}

	var SysMenusRsp []*dto.FindSysMenuRsp
	for _, v := range ens {
		cfg := a.E2DFindSysMenuRsp(v)
		SysMenusRsp = append(SysMenusRsp, cfg)
	}

	return SysMenusRsp
}
