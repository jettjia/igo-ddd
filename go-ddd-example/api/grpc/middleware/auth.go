package middleware

import (
	"context"

	middleware "github.com/jettjia/igo-pkg/pkg/xmiddleware/mgrpc"
)

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	return middleware.AuthInterceptor(ctx, []byte("reader your config.yaml"))
}
