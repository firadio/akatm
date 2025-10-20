package svc

import (
	"akatm/rpc/fams/internal/config"
	"akatm/rpc/fams/internal/repository"
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	// Repository 层
	BankCustomerRepository           repository.BankCustomerRepository
	BankAccountRepository            repository.BankAccountRepository
	UserWalletRepository             repository.UserWalletRepository
	BankAccountApplicationRepository repository.BankAccountApplicationRepository
	FundDetailRepository             repository.FundDetailRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewDB(c.Database.DSN)
	return &ServiceContext{
		Config: c,
		DB:     db,
		// 初始化 Repository
		BankCustomerRepository:           repository.NewBankCustomerRepository(db),
		BankAccountRepository:            repository.NewBankAccountRepository(db),
		UserWalletRepository:             repository.NewUserWalletRepository(db),
		BankAccountApplicationRepository: repository.NewBankAccountApplicationRepository(db),
		FundDetailRepository:             repository.NewFundDetailRepository(db),
	}
}
