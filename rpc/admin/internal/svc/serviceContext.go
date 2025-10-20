package svc

import (
	"akatm/rpc/admin/internal/config"
	"akatm/rpc/admin/internal/repository"
	"akatm/rpc/admin/orm"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	// Repository 层
	RoleRepository       repository.RoleRepository
	StaffRepository      repository.StaffRepository
	MenuRepository       repository.MenuRepository
	PermissionRepository repository.PermissionRepository
	StaffRoleRepository  repository.StaffRoleRepository
	RoleMenuRepository   repository.RoleMenuRepository
	CountryRepository    repository.CountryRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewDB(c.Database.DSN)
	return &ServiceContext{
		Config: c,
		DB:     db,
		// 初始化 Repository
		RoleRepository:       repository.NewRoleRepository(db),
		StaffRepository:      repository.NewStaffRepository(db),
		MenuRepository:       repository.NewMenuRepository(db),
		PermissionRepository: repository.NewPermissionRepository(db),
		StaffRoleRepository:  repository.NewStaffRoleRepository(db),
		RoleMenuRepository:   repository.NewRoleMenuRepository(db),
		CountryRepository:    repository.NewCountryRepository(db),
	}
}
