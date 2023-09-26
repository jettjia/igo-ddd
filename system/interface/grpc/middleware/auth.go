package middleware

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	"jettjia/go-ddd-demo-multi-common/pkg/log"
	"jettjia/go-ddd-demo-multi-system/domain/entity"
)

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	var (
		err error
	)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.NewLogger().Error("AuthInterceptor:metadata:error")
	}

	if len(md.Get("authorization")) == 0 {
		return ctx, err
	}

	token := md.Get("authorization")[0]
	token = strings.TrimPrefix(token, "Bearer ")
	claims, _ := entity.ParseToken(token) // 这里token校验已经关闭，只解析了token
	if claims == nil {
		return ctx, err
	}

	return ctx, nil
}
