package repository

import (
	"akatm/rpc/fams/orm/table"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// UserWalletRepository 用户钱包数据访问接口
type UserWalletRepository interface {
	// 创建用户钱包
	Create(wallet *table.UserWallet) error
	// 根据ID获取用户钱包
	GetByID(id uint) (*table.UserWallet, error)
	// 根据用户ID获取用户钱包
	GetByUserID(userID uint) (*table.UserWallet, error)
	// 根据用户ID和资产代号获取用户钱包
	GetByUserIDAndSymbol(userID uint, symbol string) (*table.UserWallet, error)
	// 根据客户经理ID获取用户钱包列表
	GetByManagerID(managerID uint) ([]*table.UserWallet, error)
	// 更新用户钱包
	Update(wallet *table.UserWallet) error
	// 删除用户钱包
	Delete(id uint) error
	// 获取用户钱包列表（分页）
	List(page, pageSize int64, keyword string, managerID uint, symbol string) ([]*table.UserWallet, int64, error)
	// 更新钱包余额
	UpdateBalance(walletID uint, availableBalance, frozenAmount decimal.Decimal) error
	// 冻结余额
	FreezeBalance(walletID uint, amount decimal.Decimal) error
	// 解冻余额
	UnfreezeBalance(walletID uint, amount decimal.Decimal) error
	// 扣减余额
	DeductBalance(walletID uint, amount decimal.Decimal) error
	// 增加余额
	AddBalance(walletID uint, amount decimal.Decimal) error
}

// userWalletRepository 用户钱包数据访问实现
type userWalletRepository struct {
	db *gorm.DB
}

// NewUserWalletRepository 创建用户钱包Repository
func NewUserWalletRepository(db *gorm.DB) UserWalletRepository {
	return &userWalletRepository{db: db}
}

func (r *userWalletRepository) Create(wallet *table.UserWallet) error {
	return r.db.Create(wallet).Error
}

func (r *userWalletRepository) GetByID(id uint) (*table.UserWallet, error) {
	var wallet table.UserWallet
	err := r.db.Preload("UserWalletAddresses").First(&wallet, id).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *userWalletRepository) GetByUserID(userID uint) (*table.UserWallet, error) {
	var wallet table.UserWallet
	err := r.db.Where("user_id = ?", userID).Preload("UserWalletAddresses").First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *userWalletRepository) GetByUserIDAndSymbol(userID uint, symbol string) (*table.UserWallet, error) {
	var wallet table.UserWallet
	err := r.db.Where("user_id = ? AND symbol = ?", userID, symbol).Preload("UserWalletAddresses").First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *userWalletRepository) GetByManagerID(managerID uint) ([]*table.UserWallet, error) {
	var wallets []*table.UserWallet
	err := r.db.Where("manager_id = ?", managerID).Preload("UserWalletAddresses").Find(&wallets).Error
	return wallets, err
}

func (r *userWalletRepository) Update(wallet *table.UserWallet) error {
	return r.db.Save(wallet).Error
}

func (r *userWalletRepository) Delete(id uint) error {
	return r.db.Delete(&table.UserWallet{}, id).Error
}

func (r *userWalletRepository) List(page, pageSize int64, keyword string, managerID uint, symbol string) ([]*table.UserWallet, int64, error) {
	var wallets []*table.UserWallet
	var total int64

	// 构建查询
	query := r.db.Model(&table.UserWallet{})

	// 客户经理ID过滤
	if managerID > 0 {
		query = query.Where("manager_id = ?", managerID)
	}

	// 资产代号过滤
	if symbol != "" {
		query = query.Where("symbol = ?", symbol)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Preload("UserWalletAddresses").Find(&wallets).Error; err != nil {
		return nil, 0, err
	}

	return wallets, total, nil
}

func (r *userWalletRepository) UpdateBalance(walletID uint, availableBalance, frozenAmount decimal.Decimal) error {
	return r.db.Model(&table.UserWallet{}).Where("id = ?", walletID).Updates(map[string]interface{}{
		"available_balance": availableBalance,
		"frozen_amount":     frozenAmount,
	}).Error
}

func (r *userWalletRepository) FreezeBalance(walletID uint, amount decimal.Decimal) error {
	return r.db.Model(&table.UserWallet{}).Where("id = ?", walletID).Updates(map[string]interface{}{
		"available_balance": gorm.Expr("available_balance - ?", amount),
		"frozen_amount":     gorm.Expr("frozen_amount + ?", amount),
	}).Error
}

func (r *userWalletRepository) UnfreezeBalance(walletID uint, amount decimal.Decimal) error {
	return r.db.Model(&table.UserWallet{}).Where("id = ?", walletID).Updates(map[string]interface{}{
		"available_balance": gorm.Expr("available_balance + ?", amount),
		"frozen_amount":     gorm.Expr("frozen_amount - ?", amount),
	}).Error
}

func (r *userWalletRepository) DeductBalance(walletID uint, amount decimal.Decimal) error {
	return r.db.Model(&table.UserWallet{}).Where("id = ?", walletID).Update("available_balance", gorm.Expr("available_balance - ?", amount)).Error
}

func (r *userWalletRepository) AddBalance(walletID uint, amount decimal.Decimal) error {
	return r.db.Model(&table.UserWallet{}).Where("id = ?", walletID).Update("available_balance", gorm.Expr("available_balance + ?", amount)).Error
}
