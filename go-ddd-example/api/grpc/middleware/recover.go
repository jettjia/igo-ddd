package middleware

import (
	middleware "github.com/jettjia/igo-pkg/pkg/xmiddleware/mgrpc"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

func RecoverInterceptor() grpc_recovery.Option {
	return middleware.RecoverInterceptor()
}
