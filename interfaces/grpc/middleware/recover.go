package middleware

import (
	"fmt"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jettjia/go-ddd-demo/global"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/responseutil"
	"github.com/jettjia/go-ddd-demo/infrastructure/pkg/util"
)

func RecoverInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		errorInfo := responseutil.GrpcPanic(fmt.Sprintf("%s", p)) // 全局错误信息
		global.GLog.Errorln(util.PrintJson(errorInfo.Internale))  // 错误写入日志
		return status.ErrorProto(&spb.Status{Code: int32(codes.Internal), Message: fmt.Sprintf("%v", p), Details: nil})
	})
}
