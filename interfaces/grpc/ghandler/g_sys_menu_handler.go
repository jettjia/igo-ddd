package ghandler

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	dto "github.com/jettjia/go-ddd-demo/application/dto/sys"
	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/validate"
	grpcGoodsProto "github.com/jettjia/go-ddd-demo/interfaces/grpc/proto/goods"
)

func (s *GrpcGoodsServer) CreateSysMenu(ctx context.Context, req *grpcGoodsProto.CreateSysMenuReq) (rsp *grpcGoodsProto.CreateSysMenuRsp, err error) {
	var (
		reqDto dto.CreateSysMenuReq
		rspDto *dto.CreateSysMenuRsp
	)
	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		global.GLog.Errorln("CreateSysMenu:copier.Copy:err:", err)
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

	return &grpcGoodsProto.CreateSysMenuRsp{Id: rspDto.Id}, nil
}

func (s *GrpcGoodsServer) DeleteSysMenu(ctx context.Context, req *grpcGoodsProto.DeleteSysMenuReq) (*emptypb.Empty, error) {
	var (
		reqDto dto.DelSysMenusReq
		err    error
	)

	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		global.GLog.Errorln("DeleteSysMenu:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 参数过滤
	err = validate.Validate(reqDto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	err = s.SysMenuSrv.DeleteSysMenu(ctx, &reqDto)
	if err != nil {
		return nil, status.Error(codes.NotFound, "DeleteSysMenu id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	return &emptypb.Empty{}, nil
}

func (s *GrpcGoodsServer) UpdateSysMenu(ctx context.Context, req *grpcGoodsProto.UpdateSysMenuReq) (*emptypb.Empty, error) {
	var (
		reqDto dto.UpdateSysMenuReq
		err    error
	)

	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		global.GLog.Errorln("UpdateSysMenu:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 参数过滤
	err = validate.Validate(reqDto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	err = s.SysMenuSrv.UpdateSysMenu(ctx, &reqDto)
	if err != nil {
		return nil, status.Error(codes.NotFound, "UpdateProductLog id  "+strconv.Itoa(int(req.Id))+" not exists")
	}

	return &emptypb.Empty{}, nil
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
		global.GLog.Errorln("FindSysMenuById:copier.Copy:err:", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 参数过滤
	err = validate.Validate(reqDto)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 业务处理
	if rspDto, err = s.SysMenuSrv.FindSysMenuById(ctx, &reqDto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = copier.Copy(&rspPb, &rspDto)
	if err != nil {
		global.GLog.Errorln("FindSysMenuById:rspPb:copier.Copy:err:", err)
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
		global.GLog.Errorln("FindSysMenuPage:copier.Copy:err:", err)
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
		global.GLog.Errorln("FindSysMenuPage:rspPb:copier.Copy:err:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &rspPb, nil
}
