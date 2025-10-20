package repository

import (
	"akatm/rpc/fams/orm/table"

	"gorm.io/gorm"
)

// BankAccountRepository 银行账户数据访问接口
type BankAccountRepository interface {
	// 创建银行账户
	Create(account *table.BankAccount) error
	// 根据ID获取银行账户
	GetByID(id uint) (*table.BankAccount, error)
	// 根据账号获取银行账户
	GetByAccountNumber(accountNumber string) (*table.BankAccount, error)
	// 根据客户ID获取银行账户列表
	GetByCustomerID(customerID uint) ([]*table.BankAccount, error)
	// 根据用户ID获取银行账户列表
	GetByUserID(userID uint) ([]*table.BankAccount, error)
	// 根据客户经理ID获取银行账户列表
	GetByManagerID(managerID uint) ([]*table.BankAccount, error)
	// 更新银行账户
	Update(account *table.BankAccount) error
	// 删除银行账户
	Delete(id uint) error
	// 获取银行账户列表（分页）
	List(page, pageSize int64, keyword string, managerID uint, status *int8) ([]*table.BankAccount, int64, error)
	// 检查账号是否存在
	ExistsByAccountNumber(accountNumber string) (bool, error)
	// 获取账户流水（分页）
	ListAccountTransactions(accountID uint, page, pageSize int64, transactionType, status, keyword string, startTime, endTime int64, minAmount, maxAmount string) ([]*table.FundDetail, int64, error)
}

// bankAccountRepository 银行账户数据访问实现
type bankAccountRepository struct {
	db *gorm.DB
}

// NewBankAccountRepository 创建银行账户Repository
func NewBankAccountRepository(db *gorm.DB) BankAccountRepository {
	return &bankAccountRepository{db: db}
}

func (r *bankAccountRepository) Create(account *table.BankAccount) error {
	return r.db.Create(account).Error
}

func (r *bankAccountRepository) GetByID(id uint) (*table.BankAccount, error) {
	var account table.BankAccount
	err := r.db.Preload("Customer").Preload("Manager").First(&account, id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *bankAccountRepository) GetByAccountNumber(accountNumber string) (*table.BankAccount, error) {
	var account table.BankAccount
	err := r.db.Where("account_number = ?", accountNumber).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *bankAccountRepository) GetByCustomerID(customerID uint) ([]*table.BankAccount, error) {
	var accounts []*table.BankAccount
	err := r.db.Where("customer_id = ?", customerID).Find(&accounts).Error
	return accounts, err
}

func (r *bankAccountRepository) GetByUserID(userID uint) ([]*table.BankAccount, error) {
	var accounts []*table.BankAccount
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	return accounts, err
}

func (r *bankAccountRepository) GetByManagerID(managerID uint) ([]*table.BankAccount, error) {
	var accounts []*table.BankAccount
	err := r.db.Where("manager_id = ?", managerID).Find(&accounts).Error
	return accounts, err
}

func (r *bankAccountRepository) Update(account *table.BankAccount) error {
	return r.db.Save(account).Error
}

func (r *bankAccountRepository) Delete(id uint) error {
	return r.db.Delete(&table.BankAccount{}, id).Error
}

func (r *bankAccountRepository) List(page, pageSize int64, keyword string, managerID uint, status *int8) ([]*table.BankAccount, int64, error) {
	var accounts []*table.BankAccount
	var total int64

	// 构建查询
	query := r.db.Model(&table.BankAccount{})

	// 关键词搜索
	if keyword != "" {
		query = query.Where("account_number LIKE ?", "%"+keyword+"%")
	}

	// 客户经理ID过滤
	if managerID > 0 {
		query = query.Where("manager_id = ?", managerID)
	}

	// 状态过滤
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&accounts).Error; err != nil {
		return nil, 0, err
	}

	return accounts, total, nil
}

func (r *bankAccountRepository) ExistsByAccountNumber(accountNumber string) (bool, error) {
	var count int64
	err := r.db.Model(&table.BankAccount{}).Where("account_number = ?", accountNumber).Count(&count).Error
	return count > 0, err
}

func (r *bankAccountRepository) ListAccountTransactions(accountID uint, page, pageSize int64, transactionType, status, keyword string, startTime, endTime int64, minAmount, maxAmount string) ([]*table.FundDetail, int64, error) {
	var fundDetails []*table.FundDetail
	var total int64

	// 构建查询 - 通过银行账户ID关联查询资金明细
	query := r.db.Model(&table.FundDetail{}).Where("bank_account_id = ?", accountID)

	// 交易类型筛选
	if transactionType != "" {
		query = query.Where("transaction_type = ?", transactionType)
	}

	// 状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 时间范围筛选
	if startTime > 0 {
		query = query.Where("transaction_time >= ?", startTime)
	}
	if endTime > 0 {
		query = query.Where("transaction_time <= ?", endTime)
	}

	// 金额范围筛选
	if minAmount != "" {
		query = query.Where("amount >= ?", minAmount)
	}
	if maxAmount != "" {
		query = query.Where("amount <= ?", maxAmount)
	}

	// 关键词搜索（交易单号、备注等）
	if keyword != "" {
		query = query.Where("transaction_number LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("transaction_time DESC, id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&fundDetails).Error; err != nil {
		return nil, 0, err
	}

	return fundDetails, total, nil
}
