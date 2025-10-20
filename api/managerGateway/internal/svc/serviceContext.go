package svc

import (
	"akatm/api/managerGateway/internal/config"
	"akatm/api/managerGateway/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	SignCheck rest.Middleware
	JwtAuth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	sign := middleware.NewSignCheckMiddleware(c.Signature.Salt, c.Signature.SkewSeconds)
	jwtAuth := middleware.NewJwtAuthMiddleware(func(token string) bool {
		// TODO: 实现JWT token验证逻辑
		return true // 临时返回true，后续需要实现
	})
	return &ServiceContext{
		Config:    c,
		SignCheck: sign.Handle,
		JwtAuth:   jwtAuth.Handle,
	}
}
