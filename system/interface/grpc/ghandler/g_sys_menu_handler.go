package ghandler

import (
	"context"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"jettjia/go-ddd-demo-multi-common/pkg/log"
	"jettjia/go-ddd-demo-multi-common/pkg/validate"

	"jettjia/go-ddd-demo-multi-system/application/dto"
	grpcGoodsProto "jettjia/go-ddd-demo-multi-system/interface/grpc/proto/goods"
)

func (s *GrpcGoodsServer) CreateSysMenu(ctx context.Context, req *grpcGoodsProto.CreateSysMenuReq) (rsp *grpcGoodsProto.CreateSysMenuRsp, err error) {
	var (
		reqDto dto.CreateSysMenuReq
		rspDto *dto.CreateSysMenuRsp
	)
	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Error("CreateSysMenu:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 参数过滤
	err = validate.Validate(reqDto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	if rspDto, err = s.SysMenuSrv.CreateSysMenu(ctx, &reqDto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcGoodsProto.CreateSysMenuRsp{Ulid: rspDto.Ulid}, nil
}

func (s *GrpcGoodsServer) DeleteSysMenu(ctx context.Context, req *grpcGoodsProto.DeleteSysMenuReq) (*grpcGoodsProto.Empty, error) {
	var (
		reqDto dto.DelSysMenusReq
		err    error
		rspPb  grpcGoodsProto.Empty
	)

	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Error("DeleteSysMenu:copier.Copy:err:", err)
		return &rspPb, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	err = s.SysMenuSrv.DeleteSysMenu(ctx, &reqDto)
	if err != nil {
		return &rspPb, status.Error(codes.NotFound, "DeleteSysMenu id  "+req.Ulid+" not exists")
	}

	return &rspPb, nil
}

func (s *GrpcGoodsServer) UpdateSysMenu(ctx context.Context, req *grpcGoodsProto.UpdateSysMenuReq) (*grpcGoodsProto.Empty, error) {
	var (
		reqDto dto.UpdateSysMenuReq
		err    error
		rspPb  grpcGoodsProto.Empty
	)

	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Error("UpdateSysMenu:copier.Copy:err:", err)
		return &rspPb, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	err = s.SysMenuSrv.UpdateSysMenu(ctx, &reqDto)
	if err != nil {
		return &rspPb, status.Error(codes.NotFound, "UpdateProductLog id  "+req.Ulid+" not exists")
	}

	return &rspPb, nil
}

func (s *GrpcGoodsServer) FindSysMenuById(ctx context.Context, req *grpcGoodsProto.FindSysMenuByIdReq) (*grpcGoodsProto.FindSysMenuRsp, error) {
	var (
		reqDto dto.FindSysMenuByIdReq
		rspDto *dto.FindSysMenuRsp
		err    error
		rspPb  grpcGoodsProto.FindSysMenuRsp
	)
	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Error("FindSysMenuById:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	if rspDto, err = s.SysMenuSrv.FindSysMenuById(ctx, &reqDto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = copier.Copy(&rspPb, &rspDto)
	if err != nil {
		log.NewLogger().Error("FindSysMenuById:rspPb:copier.Copy:err:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &rspPb, nil
}

func (s *GrpcGoodsServer) FindSysMenuPage(ctx context.Context, req *grpcGoodsProto.FindSysMenuPageReq) (*grpcGoodsProto.FindSysMenuPageRsp, error) {
	var (
		reqDto dto.FindSysMenuPageReq
		rspDto *dto.FindSysMenuPageRsp
		err    error
		rspPb  grpcGoodsProto.FindSysMenuPageRsp
	)
	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Error("FindSysMenuPage:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 参数过滤
	err = validate.Validate(reqDto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	if rspDto, err = s.SysMenuSrv.FindSysMenuPage(ctx, &reqDto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = copier.Copy(&rspPb, &rspDto)
	if err != nil {
		log.NewLogger().Error("FindSysMenuPage:rspPb:copier.Copy:err:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &rspPb, nil
}
