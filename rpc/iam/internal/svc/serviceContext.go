package svc

import (
	"akatm/rpc/iam/internal/config"
	"akatm/rpc/iam/internal/repository"
	"akatm/rpc/iam/orm"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	// Repository 层
	UserRepository           repository.UserRepository
	UserInviteRepository     repository.UserInviteRepository
	UserSessionRepository    repository.UserSessionRepository
	UserCredentialRepository repository.UserCredentialRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewDB(c.Database.DSN)
	return &ServiceContext{
		Config: c,
		DB:     db,
		// 初始化 Repository
		UserRepository:           repository.NewUserRepository(db),
		UserInviteRepository:     repository.NewUserInviteRepository(db),
		UserSessionRepository:    repository.NewUserSessionRepository(db),
		UserCredentialRepository: repository.NewUserCredentialRepository(db),
	}
}
