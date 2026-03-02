package gen

const TempApplicationService = `
package user

import (
	"context"

	assUser "xdata/xtext/zsvc-example/application/assembler/user"
	dtoUser "xdata/xtext/zsvc-example/application/dto/user"
	aggUser "xdata/xtext/zsvc-example/domain/aggregate/user"
	srvUser "xdata/xtext/zsvc-example/domain/srv/user"
)

type SysMenuService struct {
	sysMenuDto *assUser.SysMenuDto
	sysMenuAgg *aggUser.SysMenu
	sysMenuSvc *srvUser.SysMenu
}

func NewSysMenuService() *SysMenuService {
	return &SysMenuService{
		sysMenuDto: assUser.NewSysMenuDto(),
		sysMenuAgg: aggUser.NewSysMenuAgg(),
		sysMenuSvc: srvUser.NewSysMenuSvc(),
	}
}

func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *dtoUser.CreateSysMenuReq) (*dtoUser.CreateSysMenuRsp, error) {
	var rsp dtoUser.CreateSysMenuRsp
	en := s.sysMenuDto.D2ECreateSysMenu(req)

	ulid, err := s.sysMenuAgg.CreateSysMenu(ctx, en)
	if err != nil {
		return nil, err
	}
	rsp.Ulid = ulid

	return &rsp, nil
}

func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *dtoUser.DelSysMenusReq) error {
	en := s.sysMenuDto.D2EDeleteSysMenu(req)

	return s.sysMenuAgg.DeleteSysMenu(ctx, en)
}

func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *dtoUser.UpdateSysMenuReq) error {
	en := s.sysMenuDto.D2EUpdateSysMenu(req)

	return s.sysMenuAgg.UpdateSysMenu(ctx, en)
}

func (s *SysMenuService) FindSysMenuById(ctx context.Context, req *dtoUser.FindSysMenuByIdReq) (dto *dtoUser.FindSysMenuRsp, err error) {
	en, err := s.sysMenuSvc.FindSysMenuById(ctx, req.Ulid)
	if err != nil {
		return nil, err
	}

	dto = s.sysMenuDto.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuByQuery(ctx context.Context, req *dtoUser.FindSysMenuByQueryReq) (dto *dtoUser.FindSysMenuRsp, err error) {
	en, err := s.sysMenuSvc.FindSysMenuByQuery(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	dto = s.sysMenuDto.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuAll(ctx context.Context, req *dtoUser.FindSysMenuAllReq) (entries []*dtoUser.FindSysMenuRsp, err error) {
	ens, err := s.sysMenuSvc.FindSysMenuAll(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	entries = s.sysMenuDto.E2DGetSysMenus(ens)

	return entries, nil
}

func (s *SysMenuService) FindSysMenuPage(ctx context.Context, req *dtoUser.FindSysMenuPageReq) (*dtoUser.FindSysMenuPageRsp, error) {
	var rsp dtoUser.FindSysMenuPageRsp
	ens, pageData, err := s.sysMenuSvc.FindSysMenuPage(ctx, req.Query, req.PageData, req.SortData)
	if err != nil {
		return nil, err
	}

	entries := s.sysMenuDto.E2DGetSysMenus(ens)
	rsp.Entries = entries
	rsp.PageData = pageData

	return &rsp, nil
}

`
