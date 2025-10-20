package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// RoleRepository 角色数据访问接口
type RoleRepository interface {
	// 创建角色
	Create(role *table.Role) error
	// 根据ID获取角色
	GetByID(id uint) (*table.Role, error)
	// 根据代码获取角色
	GetByCode(code string) (*table.Role, error)
	// 更新角色
	Update(role *table.Role) error
	// 删除角色
	Delete(id uint) error
	// 获取角色列表（分页）
	List(page, pageSize int64, keyword string) ([]*table.Role, int64, error)
	// 检查角色代码是否存在
	ExistsByCode(code string) (bool, error)
}

// roleRepository 角色数据访问实现
type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository 创建角色Repository
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *table.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) GetByID(id uint) (*table.Role, error) {
	var role table.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetByCode(code string) (*table.Role, error) {
	var role table.Role
	err := r.db.Where("code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Update(role *table.Role) error {
	return r.db.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&table.Role{}, id).Error
}

func (r *roleRepository) List(page, pageSize int64, keyword string) ([]*table.Role, int64, error) {
	var roles []*table.Role
	var total int64

	// 构建查询
	query := r.db.Model(&table.Role{})

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
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&roles).Error; err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

func (r *roleRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	err := r.db.Model(&table.Role{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}
