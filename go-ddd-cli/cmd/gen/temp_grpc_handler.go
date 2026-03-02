package gen

const TempGrpcHandlerBase = `

package ghandler

import (
	grpcGoodsProto "xdata/xtext/common/proto/goods"

	sysSvc "xdata/xtext/zsvc-example/application/service"
)

type GrpcGoodsServer struct {
	grpcGoodsProto.UnimplementedGoodsServer
	SysMenuSrv *sysSvc.SysMenuService
}

`

const TempGrpcHandlerBiz = `

package ghandler

import (
	"context"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"xdata/xtext/common/pkg/validate"
	grpcGoodsProto "xdata/xtext/common/proto/goods"

	"xdata/xtext/zsvc-example/application/dto"
	"xdata/xtext/zsvc-example/infra/pkg/log"
)

func (s *GrpcGoodsServer) CreateSysMenu(ctx context.Context, req *grpcGoodsProto.CreateSysMenuReq) (rsp *grpcGoodsProto.CreateSysMenuRsp, err error) {
	var (
		reqDto dto.CreateSysMenuReq
		rspDto *dto.CreateSysMenuRsp
	)
	// 参数解析
	if err = copier.Copy(&reqDto, &req); err != nil {
		log.NewLogger().Errorln("CreateSysMenu:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("DeleteSysMenu:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("UpdateSysMenu:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("FindSysMenuById:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("FindSysMenuById:rspPb:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("FindSysMenuPage:copier.Copy:err:", err)
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
		log.NewLogger().Errorln("FindSysMenuPage:rspPb:copier.Copy:err:", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &rspPb, nil
}
`
