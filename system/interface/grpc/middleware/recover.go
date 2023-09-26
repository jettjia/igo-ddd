package middleware

import (
	"fmt"
	"jettjia/go-ddd-demo-multi-common/pkg/log"
	"jettjia/go-ddd-demo-multi-common/pkg/response"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoverInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		errorInfo := response.GrpcPanic(err) // 全局错误信息
		log.NewLogger().WithFields(logrus.Fields{"error_detail": errorInfo.Internal}).Errorln(fmt.Sprintf("%v", p))
		return status.ErrorProto(&spb.Status{Code: int32(codes.Internal), Message: fmt.Sprintf("%v", p), Details: nil})
	})
}
