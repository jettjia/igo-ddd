package service

import (
	"context"

	"jettjia/go-ddd-demo-multi-system/application/assembler"
	"jettjia/go-ddd-demo-multi-system/application/dto"
	"jettjia/go-ddd-demo-multi-system/domain/aggregate"
	"jettjia/go-ddd-demo-multi-system/domain/srv"
)

type SysMenuService struct {
	sysMenuReq *assembler.SysMenuReq
	sysMenuRsp *assembler.SysMenuRsp
	sysMenuAgg *aggregate.SysMenu
	sysMenuSvc *srv.SysMenu
}

func NewSysMenuService(sysMenuReq *assembler.SysMenuReq, sysMenuRsp *assembler.SysMenuRsp, sysMenuAgg *aggregate.SysMenu, sysMenuSvc *srv.SysMenu) *SysMenuService {
	return &SysMenuService{
		sysMenuReq: sysMenuReq,
		sysMenuRsp: sysMenuRsp,
		sysMenuAgg: sysMenuAgg,
		sysMenuSvc: sysMenuSvc,
	}
}

func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *dto.CreateSysMenuReq) (*dto.CreateSysMenuRsp, error) {
	var rsp dto.CreateSysMenuRsp
	en := s.sysMenuReq.D2ECreateSysMenu(req)

	ulid, err := s.sysMenuAgg.CreateSysMenu(ctx, en)
	if err != nil {
		return nil, err
	}
	rsp.Ulid = ulid

	return &rsp, nil
}

func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *dto.DelSysMenusReq) error {
	en := s.sysMenuReq.D2EDeleteSysMenu(req)

	return s.sysMenuAgg.DeleteSysMenu(ctx, en)
}

func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *dto.UpdateSysMenuReq) error {
	en := s.sysMenuReq.D2EUpdateSysMenu(req)

	return s.sysMenuAgg.UpdateSysMenu(ctx, en)
}

func (s *SysMenuService) FindSysMenuById(ctx context.Context, req *dto.FindSysMenuByIdReq) (dto *dto.FindSysMenuRsp, err error) {
	en, err := s.sysMenuSvc.FindSysMenuById(ctx, req.Ulid)
	if err != nil {
		return nil, err
	}

	dto = s.sysMenuRsp.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuByQuery(ctx context.Context, req *dto.FindSysMenuByQueryReq) (dto *dto.FindSysMenuRsp, err error) {
	en, err := s.sysMenuSvc.FindSysMenuByQuery(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	dto = s.sysMenuRsp.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuAll(ctx context.Context, req *dto.FindSysMenuAllReq) (entries []*dto.FindSysMenuRsp, err error) {
	ens, err := s.sysMenuSvc.FindSysMenuAll(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	entries = s.sysMenuRsp.E2DGetSysMenus(ens)

	return entries, nil
}

func (s *SysMenuService) FindSysMenuPage(ctx context.Context, req *dto.FindSysMenuPageReq) (*dto.FindSysMenuPageRsp, error) {
	var rsp dto.FindSysMenuPageRsp
	ens, pageData, err := s.sysMenuSvc.FindSysMenuPage(ctx, req.Query, req.PageData, req.SortData)
	if err != nil {
		return nil, err
	}

	entries := s.sysMenuRsp.E2DGetSysMenus(ens)
	rsp.Entries = entries
	rsp.PageData = pageData

	return &rsp, nil
}
