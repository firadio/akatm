package repository

import (
	"akatm/rpc/fams/orm/table"

	"gorm.io/gorm"
)

// BankCustomerRepository 银行客户数据访问接口
type BankCustomerRepository interface {
	// 创建银行客户
	Create(customer *table.BankCustomer) error
	// 根据ID获取银行客户
	GetByID(id uint) (*table.BankCustomer, error)
	// 根据用户ID获取银行客户
	GetByUserID(userID uint) (*table.BankCustomer, error)
	// 根据客户经理ID获取银行客户列表
	GetByManagerID(managerID uint) ([]*table.BankCustomer, error)
	// 更新银行客户
	Update(customer *table.BankCustomer) error
	// 删除银行客户
	Delete(id uint) error
	// 获取银行客户列表（分页）
	List(page, pageSize int64, keyword string, managerID uint) ([]*table.BankCustomer, int64, error)
	// 检查手机号是否存在
	ExistsByPhone(phone string) (bool, error)
	// 检查邮箱是否存在
	ExistsByEmail(email string) (bool, error)
}

// bankCustomerRepository 银行客户数据访问实现
type bankCustomerRepository struct {
	db *gorm.DB
}

// NewBankCustomerRepository 创建银行客户Repository
func NewBankCustomerRepository(db *gorm.DB) BankCustomerRepository {
	return &bankCustomerRepository{db: db}
}

func (r *bankCustomerRepository) Create(customer *table.BankCustomer) error {
	return r.db.Create(customer).Error
}

func (r *bankCustomerRepository) GetByID(id uint) (*table.BankCustomer, error) {
	var customer table.BankCustomer
	err := r.db.Preload("BankAccounts").First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *bankCustomerRepository) GetByUserID(userID uint) (*table.BankCustomer, error) {
	var customer table.BankCustomer
	err := r.db.Where("user_id = ?", userID).Preload("BankAccounts").First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *bankCustomerRepository) GetByManagerID(managerID uint) ([]*table.BankCustomer, error) {
	var customers []*table.BankCustomer
	err := r.db.Where("manager_id = ?", managerID).Preload("BankAccounts").Find(&customers).Error
	return customers, err
}

func (r *bankCustomerRepository) Update(customer *table.BankCustomer) error {
	return r.db.Save(customer).Error
}

func (r *bankCustomerRepository) Delete(id uint) error {
	return r.db.Delete(&table.BankCustomer{}, id).Error
}

func (r *bankCustomerRepository) List(page, pageSize int64, keyword string, managerID uint) ([]*table.BankCustomer, int64, error) {
	var customers []*table.BankCustomer
	var total int64

	// 构建查询
	query := r.db.Model(&table.BankCustomer{})

	// 关键词搜索
	if keyword != "" {
		query = query.Where("first_name LIKE ? OR last_name LIKE ? OR phone LIKE ? OR email LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 客户经理ID过滤
	if managerID > 0 {
		query = query.Where("manager_id = ?", managerID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Preload("BankAccounts").Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (r *bankCustomerRepository) ExistsByPhone(phone string) (bool, error) {
	var count int64
	err := r.db.Model(&table.BankCustomer{}).Where("phone = ?", phone).Count(&count).Error
	return count > 0, err
}

func (r *bankCustomerRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&table.BankCustomer{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
