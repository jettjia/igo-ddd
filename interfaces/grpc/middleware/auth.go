package middleware

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	entity "github.com/jettjia/go-ddd-demo/domain/entity/jwt"
	"github.com/jettjia/go-ddd-demo/global"
)

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	var (
		err error
	)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		global.GLog.Errorln("AuthInterceptor:metadata:error")
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
	global.GCustomerInfo = claims.CustomerInfo

	return ctx, nil
}
