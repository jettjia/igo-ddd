package sys

import (
	"context"

	assembler "github.com/jettjia/go-ddd-demo/application/assembler/sys"
	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	agg "github.com/jettjia/go-ddd-demo/domain/aggregate/sys"
	svc "github.com/jettjia/go-ddd-demo/domain/service/sys"
)

type SysMenuService struct {
	sysMenuReq *assembler.SysMenuReq
	sysMenuRsp *assembler.SysMenuRsp
	sysMenuAgg agg.SysMenuAgg
	sysMenuSvc svc.SysMenuSvc
}

func NewSysMenuService() *SysMenuService {
	return &SysMenuService{
		sysMenuReq: assembler.NewSysMenuReq(),
		sysMenuRsp: assembler.NewSysMenuRsp(),
		sysMenuAgg: agg.NewSysMenuAgg(),
		sysMenuSvc: svc.NewSysMenuSvc(),
	}
}

func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *dto.CreateSysMenuReq) (*dto.CreateSysMenuRsp, error) {
	var rsp dto.CreateSysMenuRsp
	en := s.sysMenuReq.D2ECreateSysMenu(req)

	id, err := s.sysMenuAgg.CreateSysMenu(ctx, en)
	if err != nil {
		return nil, err
	}
	rsp.Id = id

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
	en, err := s.sysMenuSvc.FindSysMenuById(ctx, req.Id)
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
