package user

import (
	"context"

	assUser "github.com/jettjia/ddddemo/application/assembler/user"
	dtoUser "github.com/jettjia/ddddemo/application/dto/user"
	aggUser "github.com/jettjia/ddddemo/domain/aggregate/user"
	srvUser "github.com/jettjia/ddddemo/domain/srv/user"
)

type SysUserService struct {
	sysUserDto *assUser.SysUserDto
	sysUserAgg *aggUser.SysUser
	sysUserSvc *srvUser.SysUser
}

func NewSysUserService() *SysUserService {
	return &SysUserService{
		sysUserDto: assUser.NewSysUserDto(),
		sysUserAgg: aggUser.NewSysUserAgg(),
		sysUserSvc: srvUser.NewSysUserSvc(),
	}
}

func (s *SysUserService) CreateSysUser(ctx context.Context, req *dtoUser.CreateSysUserReq) (*dtoUser.CreateSysUserRsp, error) {
	var rsp dtoUser.CreateSysUserRsp
	en := s.sysUserDto.D2ECreateSysUser(req)

	ulid, err := s.sysUserAgg.CreateSysUser(ctx, en)
	if err != nil {
		return nil, err
	}
	rsp.Ulid = ulid

	return &rsp, nil
}

func (s *SysUserService) DeleteSysUser(ctx context.Context, req *dtoUser.DelSysUsersReq) error {
	en := s.sysUserDto.D2EDeleteSysUser(req)

	return s.sysUserAgg.DeleteSysUser(ctx, en)
}

func (s *SysUserService) UpdateSysUser(ctx context.Context, req *dtoUser.UpdateSysUserReq) error {
	en := s.sysUserDto.D2EUpdateSysUser(req)

	return s.sysUserAgg.UpdateSysUser(ctx, en)
}

func (s *SysUserService) FindSysUserById(ctx context.Context, req *dtoUser.FindSysUserByIdReq) (dto *dtoUser.FindSysUserRsp, err error) {
	en, err := s.sysUserSvc.FindSysUserById(ctx, req.Ulid)
	if err != nil {
		return nil, err
	}

	dto = s.sysUserDto.E2DFindSysUserRsp(en)

	return dto, nil
}

func (s *SysUserService) FindSysUserByQuery(ctx context.Context, req *dtoUser.FindSysUserByQueryReq) (dto *dtoUser.FindSysUserRsp, err error) {
	en, err := s.sysUserSvc.FindSysUserByQuery(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	dto = s.sysUserDto.E2DFindSysUserRsp(en)

	return dto, nil
}

func (s *SysUserService) FindSysUserAll(ctx context.Context, req *dtoUser.FindSysUserAllReq) (entries []*dtoUser.FindSysUserRsp, err error) {
	ens, err := s.sysUserSvc.FindSysUserAll(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	entries = s.sysUserDto.E2DGetSysUsers(ens)

	return entries, nil
}

func (s *SysUserService) FindSysUserPage(ctx context.Context, req *dtoUser.FindSysUserPageReq) (*dtoUser.FindSysUserPageRsp, error) {
	var rsp dtoUser.FindSysUserPageRsp
	ens, pageData, err := s.sysUserSvc.FindSysUserPage(ctx, req.Query, req.PageData, req.SortData)
	if err != nil {
		return nil, err
	}

	entries := s.sysUserDto.E2DGetSysUsers(ens)
	rsp.Entries = entries
	rsp.PageData = pageData

	return &rsp, nil
}
