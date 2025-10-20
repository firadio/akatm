package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// MenuRepository 菜单数据访问接口
type MenuRepository interface {
	// 创建菜单
	Create(menu *table.Menu) error
	// 根据ID获取菜单
	GetByID(id uint) (*table.Menu, error)
	// 更新菜单
	Update(menu *table.Menu) error
	// 删除菜单
	Delete(id uint) error
	// 获取菜单列表（分页）
	List(page, pageSize int64, keyword string, menuType int8) ([]*table.Menu, int64, error)
	// 获取所有菜单（树形结构）
	GetTree(menuType int8) ([]*table.Menu, error)
}

// menuRepository 菜单数据访问实现
type menuRepository struct {
	db *gorm.DB
}

// NewMenuRepository 创建菜单Repository
func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(menu *table.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuRepository) GetByID(id uint) (*table.Menu, error) {
	var menu table.Menu
	err := r.db.First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Update(menu *table.Menu) error {
	return r.db.Save(menu).Error
}

func (r *menuRepository) Delete(id uint) error {
	return r.db.Delete(&table.Menu{}, id).Error
}

func (r *menuRepository) List(page, pageSize int64, keyword string, menuType int8) ([]*table.Menu, int64, error) {
	var menus []*table.Menu
	var total int64

	// 构建查询
	query := r.db.Model(&table.Menu{})

	// 类型筛选
	if menuType > 0 {
		query = query.Where("type = ?", menuType)
	}

	// 关键词搜索
	if keyword != "" {
		query = query.Where("name LIKE ? OR path LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, total, nil
}

func (r *menuRepository) GetTree(menuType int8) ([]*table.Menu, error) {
	var menus []*table.Menu
	query := r.db.Where("parent_id = 0")

	// 类型筛选
	if menuType > 0 {
		query = query.Where("type = ?", menuType)
	}

	err := query.Preload("Children").Find(&menus).Error
	return menus, err
}
