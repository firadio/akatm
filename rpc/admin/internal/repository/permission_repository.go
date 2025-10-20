package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// PermissionRepository 权限数据访问接口
type PermissionRepository interface {
	// 创建权限
	Create(permission *table.Permission) error
	// 根据ID获取权限
	GetByID(id uint) (*table.Permission, error)
	// 根据代码获取权限
	GetByCode(code string) (*table.Permission, error)
	// 更新权限
	Update(permission *table.Permission) error
	// 删除权限
	Delete(id uint) error
	// 获取权限列表（分页）
	List(page, pageSize int64, keyword string) ([]*table.Permission, int64, error)
	// 检查权限代码是否存在
	ExistsByCode(code string) (bool, error)
}

// permissionRepository 权限数据访问实现
type permissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository 创建权限Repository
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(permission *table.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) GetByID(id uint) (*table.Permission, error) {
	var permission table.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) GetByCode(code string) (*table.Permission, error) {
	var permission table.Permission
	err := r.db.Where("code = ?", code).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) Update(permission *table.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return r.db.Delete(&table.Permission{}, id).Error
}

func (r *permissionRepository) List(page, pageSize int64, keyword string) ([]*table.Permission, int64, error) {
	var permissions []*table.Permission
	var total int64

	// 构建查询
	query := r.db.Model(&table.Permission{})

	// 关键词搜索
	if keyword != "" {
		query = query.Where("name LIKE ? OR code LIKE ? OR description LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&permissions).Error; err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

func (r *permissionRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	err := r.db.Model(&table.Permission{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}
