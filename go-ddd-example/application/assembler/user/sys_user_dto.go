package user

import (
	"github.com/jinzhu/copier"

	dtoUser "github.com/jettjia/ddddemo/application/dto/user"
	entityUser "github.com/jettjia/ddddemo/domain/entity/user"
)

// SysUserDto 请求参数
type SysUserDto struct {
}

// NewSysUserDto NewUSysUserReq
func NewSysUserDto() *SysUserDto {
	return &SysUserDto{}
}

//////////////////////////////////////////////////////////////////
// dto to entity

// D2ECreateSysUser dto转换成entity
func (a *SysUserDto) D2ECreateSysUser(dto *dtoUser.CreateSysUserReq) *entityUser.SysUser {
	var rspEn entityUser.SysUser

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}

	return &rspEn
}

// D2EDeleteSysUser dto转换成entity
func (a *SysUserDto) D2EDeleteSysUser(dto *dtoUser.DelSysUsersReq) *entityUser.SysUser {
	var rspEn entityUser.SysUser

	rspEn.Ulid = dto.Ulid

	return &rspEn
}

// D2EUpdateSysUser dto转换成entity
func (a *SysUserDto) D2EUpdateSysUser(dto *dtoUser.UpdateSysUserReq) *entityUser.SysUser {
	var rspEn entityUser.SysUser

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}
	return &rspEn
}

// D2ELoginSysUser dto转换成entity 登录
func (a *SysUserDto) D2ELoginSysUser(dto *dtoUser.LoginReq) *entityUser.SysUser {
	var rspEn entityUser.SysUser

	if err := copier.Copy(&rspEn, &dto); err != nil {
		panic(any(err))
	}
	return &rspEn
}

// ////////////////////////////////////////////////////////////////
// entity to dto

// E2DCreateSysUser dto转换成entity
func (a *SysUserDto) E2DCreateSysUser(en *entityUser.SysUser) (dto *dtoUser.CreateSysUserRsp) {
	if en == nil {
		return &dtoUser.CreateSysUserRsp{}
	}
	return &dtoUser.CreateSysUserRsp{Ulid: en.Ulid}
}

// E2DFindSysUserRsp entity转换成dto
func (a *SysUserDto) E2DFindSysUserRsp(en *entityUser.SysUser) *dtoUser.FindSysUserRsp {
	var rspDto dtoUser.FindSysUserRsp

	if err := copier.Copy(&rspDto, &en); err != nil {
		panic(any(err))
	}

	return &rspDto
}

// E2DGetSysUsers entity转换成dto
func (a *SysUserDto) E2DGetSysUsers(ens []*entityUser.SysUser) []*dtoUser.FindSysUserRsp {
	if len(ens) == 0 {
		return []*dtoUser.FindSysUserRsp{}
	}

	var SysUsersRsp []*dtoUser.FindSysUserRsp
	for _, v := range ens {
		cfg := a.E2DFindSysUserRsp(v)
		SysUsersRsp = append(SysUsersRsp, cfg)
	}

	return SysUsersRsp
}
