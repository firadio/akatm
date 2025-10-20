package svc

import (
	"akatm/rpc/mail/internal/config"
	"akatm/rpc/mail/internal/repository"
	"akatm/rpc/mail/orm"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	// Repository 层
	EmailVerifyRepository repository.EmailVerifyRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewDB(c.Database.DSN)
	return &ServiceContext{
		Config: c,
		DB:     db,
		// 初始化 Repository
		EmailVerifyRepository: repository.NewEmailVerifyRepository(db),
	}
}
