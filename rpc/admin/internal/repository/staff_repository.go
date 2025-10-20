package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// StaffRepository 员工数据访问接口
type StaffRepository interface {
	// 创建员工
	Create(staff *table.Staff) error
	// 根据ID获取员工
	GetByID(id uint) (*table.Staff, error)
	// 根据邮箱获取员工
	GetByEmail(email string) (*table.Staff, error)
	// 更新员工
	Update(staff *table.Staff) error
	// 删除员工
	Delete(id uint) error
	// 获取员工列表（分页）
	List(page, pageSize int64, keyword string) ([]*table.Staff, int64, error)
	// 检查邮箱是否存在
	ExistsByEmail(email string) (bool, error)
}

// staffRepository 员工数据访问实现
type staffRepository struct {
	db *gorm.DB
}

// NewStaffRepository 创建员工Repository
func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) Create(staff *table.Staff) error {
	return r.db.Create(staff).Error
}

func (r *staffRepository) GetByID(id uint) (*table.Staff, error) {
	var staff table.Staff
	err := r.db.First(&staff, id).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) GetByEmail(email string) (*table.Staff, error) {
	var staff table.Staff
	err := r.db.Where("email = ?", email).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) Update(staff *table.Staff) error {
	return r.db.Save(staff).Error
}

func (r *staffRepository) Delete(id uint) error {
	return r.db.Delete(&table.Staff{}, id).Error
}

func (r *staffRepository) List(page, pageSize int64, keyword string) ([]*table.Staff, int64, error) {
	var staffs []*table.Staff
	var total int64

	// 构建查询
	query := r.db.Model(&table.Staff{})

	// 关键词搜索
	if keyword != "" {
		query = query.Where("name LIKE ? OR email LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&staffs).Error; err != nil {
		return nil, 0, err
	}

	return staffs, total, nil
}

func (r *staffRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&table.Staff{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
