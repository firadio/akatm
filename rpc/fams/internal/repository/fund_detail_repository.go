package repository

import (
	"akatm/rpc/fams/orm/table"

	"gorm.io/gorm"
)

// FundDetailRepository 资金明细数据访问接口
type FundDetailRepository interface {
	// 创建资金明细记录
	Create(fundDetail *table.FundDetail) error
	// 根据ID获取资金明细
	GetByID(id uint) (*table.FundDetail, error)
	// 获取资金明细列表（分页）
	List(page, pageSize int64, keyword, userType, transactionType, status, currency string, startTime, endTime int64, minAmount, maxAmount string) ([]*table.FundDetail, int64, error)
	// 获取用户资金明细
	GetUserFundDetails(userId uint, page, pageSize int64, transactionType, status string, startTime, endTime int64) ([]*table.FundDetail, int64, error)
	// 获取代理及其下级用户资金明细
	GetAgentFundDetails(agentId uint, includeSubAgents bool, page, pageSize int64, transactionType, status string, startTime, endTime int64) ([]*table.FundDetail, int64, error)
	// 获取资金统计总览
	GetFundSummary(startTime, endTime int64, userType, currency string) (*FundSummaryData, error)
}

// fundDetailRepository 资金明细数据访问实现
type fundDetailRepository struct {
	db *gorm.DB
}

// FundSummaryData 资金统计总览数据
type FundSummaryData struct {
	TotalDeposit     string
	TotalWithdrawal  string
	TotalFee         string
	ActiveUsers      int64
	TransactionCount int64
	CurrencyStats    []CurrencyStat
}

// CurrencyStat 币种统计
type CurrencyStat struct {
	Currency         string
	DepositAmount    string
	WithdrawalAmount string
	FeeAmount        string
	TransactionCount int64
}

// NewFundDetailRepository 创建资金明细Repository
func NewFundDetailRepository(db *gorm.DB) FundDetailRepository {
	return &fundDetailRepository{db: db}
}

func (r *fundDetailRepository) Create(fundDetail *table.FundDetail) error {
	return r.db.Create(fundDetail).Error
}

func (r *fundDetailRepository) GetByID(id uint) (*table.FundDetail, error) {
	var fundDetail table.FundDetail
	err := r.db.First(&fundDetail, id).Error
	if err != nil {
		return nil, err
	}
	return &fundDetail, nil
}

func (r *fundDetailRepository) List(page, pageSize int64, keyword, userType, transactionType, status, currency string, startTime, endTime int64, minAmount, maxAmount string) ([]*table.FundDetail, int64, error) {
	var fundDetails []*table.FundDetail
	var total int64

	// 构建查询
	query := r.db.Model(&table.FundDetail{})

	// 用户类型筛选
	if userType != "" {
		query = query.Where("user_type = ?", userType)
	}

	// 交易类型筛选
	if transactionType != "" {
		query = query.Where("transaction_type = ?", transactionType)
	}

	// 状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 币种筛选
	if currency != "" {
		query = query.Where("currency = ?", currency)
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

	// 关键词搜索（交易单号）
	if keyword != "" {
		query = query.Where("transaction_number LIKE ?", "%"+keyword+"%")
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

func (r *fundDetailRepository) GetUserFundDetails(userId uint, page, pageSize int64, transactionType, status string, startTime, endTime int64) ([]*table.FundDetail, int64, error) {
	var fundDetails []*table.FundDetail
	var total int64

	// 构建查询
	query := r.db.Model(&table.FundDetail{}).Where("user_id = ?", userId)

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

func (r *fundDetailRepository) GetAgentFundDetails(agentId uint, includeSubAgents bool, page, pageSize int64, transactionType, status string, startTime, endTime int64) ([]*table.FundDetail, int64, error) {
	var fundDetails []*table.FundDetail
	var total int64

	// 构建查询
	query := r.db.Model(&table.FundDetail{})

	if includeSubAgents {
		// 包含下级代理：查询代理及其所有下级用户的资金明细
		// 这里需要根据用户层级关系来查询，暂时简化处理
		query = query.Where("parent_user_id = ? OR user_id = ?", agentId, agentId)
	} else {
		// 只查询代理直接管理的用户
		query = query.Where("parent_user_id = ?", agentId)
	}

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

func (r *fundDetailRepository) GetFundSummary(startTime, endTime int64, userType, currency string) (*FundSummaryData, error) {
	var summary FundSummaryData

	// 构建基础查询
	query := r.db.Model(&table.FundDetail{})

	// 用户类型筛选
	if userType != "" {
		query = query.Where("user_type = ?", userType)
	}

	// 币种筛选
	if currency != "" {
		query = query.Where("currency = ?", currency)
	}

	// 时间范围筛选
	if startTime > 0 {
		query = query.Where("transaction_time >= ?", startTime)
	}
	if endTime > 0 {
		query = query.Where("transaction_time <= ?", endTime)
	}

	// 统计总存款
	var totalDeposit string
	err := query.Where("transaction_type = ? AND status = ?", "deposit", "completed").Select("SUM(amount)").Scan(&totalDeposit).Error
	if err != nil {
		return nil, err
	}
	summary.TotalDeposit = totalDeposit

	// 统计总提现
	var totalWithdrawal string
	err = query.Where("transaction_type = ? AND status = ?", "withdrawal", "completed").Select("SUM(amount)").Scan(&totalWithdrawal).Error
	if err != nil {
		return nil, err
	}
	summary.TotalWithdrawal = totalWithdrawal

	// 统计总手续费
	var totalFee string
	err = query.Select("SUM(fee)").Scan(&totalFee).Error
	if err != nil {
		return nil, err
	}
	summary.TotalFee = totalFee

	// 统计活跃用户数
	err = query.Select("COUNT(DISTINCT user_id)").Scan(&summary.ActiveUsers).Error
	if err != nil {
		return nil, err
	}

	// 统计交易笔数
	err = query.Count(&summary.TransactionCount).Error
	if err != nil {
		return nil, err
	}

	// 按币种统计
	var currencyStats []CurrencyStat
	err = query.Select("currency, SUM(CASE WHEN transaction_type = 'deposit' AND status = 'completed' THEN amount ELSE 0 END) as deposit_amount, SUM(CASE WHEN transaction_type = 'withdrawal' AND status = 'completed' THEN amount ELSE 0 END) as withdrawal_amount, SUM(fee) as fee_amount, COUNT(*) as transaction_count").Group("currency").Scan(&currencyStats).Error
	if err != nil {
		return nil, err
	}
	summary.CurrencyStats = currencyStats

	return &summary, nil
}
