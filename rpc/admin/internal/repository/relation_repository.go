package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// StaffRoleRepository 员工角色关联数据访问接口
type StaffRoleRepository interface {
	// 创建员工角色关联
	Create(staffRole *table.StaffRole) error
	// 删除员工的所有角色关联
	DeleteByStaffID(staffID uint) error
	// 删除特定的员工角色关联
	DeleteByStaffIDAndRoleID(staffID, roleID uint) error
	// 检查员工角色关联是否存在
	ExistsByStaffIDAndRoleID(staffID, roleID uint) (bool, error)
	// 获取员工的所有角色
	GetRolesByStaffID(staffID uint) ([]*table.Role, error)
	// 获取角色的所有员工
	GetStaffsByRoleID(roleID uint) ([]*table.Staff, error)
}

// staffRoleRepository 员工角色关联数据访问实现
type staffRoleRepository struct {
	db *gorm.DB
}

// NewStaffRoleRepository 创建员工角色关联Repository
func NewStaffRoleRepository(db *gorm.DB) StaffRoleRepository {
	return &staffRoleRepository{db: db}
}

func (r *staffRoleRepository) Create(staffRole *table.StaffRole) error {
	return r.db.Create(staffRole).Error
}

func (r *staffRoleRepository) DeleteByStaffID(staffID uint) error {
	return r.db.Where("staff_id = ?", staffID).Delete(&table.StaffRole{}).Error
}

func (r *staffRoleRepository) DeleteByStaffIDAndRoleID(staffID, roleID uint) error {
	return r.db.Where("staff_id = ? AND role_id = ?", staffID, roleID).Delete(&table.StaffRole{}).Error
}

func (r *staffRoleRepository) ExistsByStaffIDAndRoleID(staffID, roleID uint) (bool, error) {
	var count int64
	err := r.db.Model(&table.StaffRole{}).Where("staff_id = ? AND role_id = ?", staffID, roleID).Count(&count).Error
	return count > 0, err
}

func (r *staffRoleRepository) GetRolesByStaffID(staffID uint) ([]*table.Role, error) {
	var roles []*table.Role
	err := r.db.Table("roles").
		Joins("JOIN staff_roles ON roles.id = staff_roles.role_id").
		Where("staff_roles.staff_id = ?", staffID).
		Find(&roles).Error
	return roles, err
}

func (r *staffRoleRepository) GetStaffsByRoleID(roleID uint) ([]*table.Staff, error) {
	var staffs []*table.Staff
	err := r.db.Table("staffs").
		Joins("JOIN staff_roles ON staffs.id = staff_roles.staff_id").
		Where("staff_roles.role_id = ?", roleID).
		Find(&staffs).Error
	return staffs, err
}

// RoleMenuRepository 角色菜单关联数据访问接口
type RoleMenuRepository interface {
	// 创建角色菜单关联
	Create(roleMenu *table.RoleMenu) error
	// 删除角色的所有菜单关联
	DeleteByRoleID(roleID uint) error
	// 删除特定的角色菜单关联
	DeleteByRoleIDAndMenuID(roleID, menuID uint) error
	// 检查角色菜单关联是否存在
	ExistsByRoleIDAndMenuID(roleID, menuID uint) (bool, error)
	// 获取角色的所有菜单
	GetMenusByRoleID(roleID uint) ([]*table.Menu, error)
	// 获取菜单的所有角色
	GetRolesByMenuID(menuID uint) ([]*table.Role, error)
}

// roleMenuRepository 角色菜单关联数据访问实现
type roleMenuRepository struct {
	db *gorm.DB
}

// NewRoleMenuRepository 创建角色菜单关联Repository
func NewRoleMenuRepository(db *gorm.DB) RoleMenuRepository {
	return &roleMenuRepository{db: db}
}

func (r *roleMenuRepository) Create(roleMenu *table.RoleMenu) error {
	return r.db.Create(roleMenu).Error
}

func (r *roleMenuRepository) DeleteByRoleID(roleID uint) error {
	return r.db.Where("role_id = ?", roleID).Delete(&table.RoleMenu{}).Error
}

func (r *roleMenuRepository) DeleteByRoleIDAndMenuID(roleID, menuID uint) error {
	return r.db.Where("role_id = ? AND menu_id = ?", roleID, menuID).Delete(&table.RoleMenu{}).Error
}

func (r *roleMenuRepository) ExistsByRoleIDAndMenuID(roleID, menuID uint) (bool, error) {
	var count int64
	err := r.db.Model(&table.RoleMenu{}).Where("role_id = ? AND menu_id = ?", roleID, menuID).Count(&count).Error
	return count > 0, err
}

func (r *roleMenuRepository) GetMenusByRoleID(roleID uint) ([]*table.Menu, error) {
	var menus []*table.Menu
	err := r.db.Table("menus").
		Joins("JOIN role_menus ON menus.id = role_menus.menu_id").
		Where("role_menus.role_id = ?", roleID).
		Find(&menus).Error
	return menus, err
}

func (r *roleMenuRepository) GetRolesByMenuID(menuID uint) ([]*table.Role, error) {
	var roles []*table.Role
	err := r.db.Table("roles").
		Joins("JOIN role_menus ON roles.id = role_menus.role_id").
		Where("role_menus.menu_id = ?", menuID).
		Find(&roles).Error
	return roles, err
}
