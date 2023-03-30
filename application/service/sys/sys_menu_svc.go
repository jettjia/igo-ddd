package service

import (
	"context"

	assembler "github.com/jettjia/go-ddd-demo/application/assembler/sys"
	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	dAggregate "github.com/jettjia/go-ddd-demo/domain/aggregate/sys"
)

type SysMenuService struct {
	SysMenuReq *assembler.SysMenuReq
	SysMenuRsp *assembler.SysMenuRsp
	SysMenuAgg dAggregate.SysMenuAgg
}

func NewSysMenuService() *SysMenuService {
	return &SysMenuService{
		SysMenuReq: assembler.NewSysMenuReq(),
		SysMenuRsp: assembler.NewSysMenuRsp(),
		SysMenuAgg: dAggregate.NewSysMenuAgg(),
	}
}

func (s *SysMenuService) CreateSysMenu(ctx context.Context, req *dto.CreateSysMenuReq) (*dto.CreateSysMenuRsp, error) {
	var rsp dto.CreateSysMenuRsp
	en := s.SysMenuReq.D2ECreateSysMenu(req)

	id, err := s.SysMenuAgg.CreateSysMenu(ctx, en)
	if err != nil {
		return nil, err
	}
	rsp.Id = id

	return &rsp, nil
}

func (s *SysMenuService) DeleteSysMenu(ctx context.Context, req *dto.DelSysMenusReq) error {
	en := s.SysMenuReq.D2EDeleteSysMenu(req)

	return s.SysMenuAgg.DeleteSysMenu(ctx, en)
}

func (s *SysMenuService) UpdateSysMenu(ctx context.Context, req *dto.UpdateSysMenuReq) error {
	en := s.SysMenuReq.D2EUpdateSysMenu(req)

	return s.SysMenuAgg.UpdateSysMenu(ctx, en)
}

func (s *SysMenuService) FindSysMenuById(ctx context.Context, req *dto.FindSysMenuByIdReq) (dto *dto.FindSysMenuRsp, err error) {
	en, err := s.SysMenuAgg.FindSysMenuById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dto = s.SysMenuRsp.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuByQuery(ctx context.Context, req *dto.FindSysMenuByQueryReq) (dto *dto.FindSysMenuRsp, err error) {
	en, err := s.SysMenuAgg.FindSysMenuByQuery(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	dto = s.SysMenuRsp.E2DFindSysMenuRsp(en)

	return dto, nil
}

func (s *SysMenuService) FindSysMenuAll(ctx context.Context, req *dto.FindSysMenuAllReq) (entries []*dto.FindSysMenuRsp, err error) {
	ens, err := s.SysMenuAgg.FindSysMenuAll(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	entries = s.SysMenuRsp.E2DGetSysMenus(ens)

	return entries, nil
}

func (s *SysMenuService) FindSysMenuPage(ctx context.Context, req *dto.FindSysMenuPageReq) (*dto.FindSysMenuPageRsp, error) {
	var rsp dto.FindSysMenuPageRsp
	ens, pageData, err := s.SysMenuAgg.FindSysMenuPage(ctx, req.Query, req.PageData, req.SortData)
	if err != nil {
		return nil, err
	}

	entries := s.SysMenuRsp.E2DGetSysMenus(ens)
	rsp.Entries = entries
	rsp.PageData = pageData

	return &rsp, nil
}
