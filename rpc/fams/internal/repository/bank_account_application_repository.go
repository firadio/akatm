package repository

import (
	"akatm/rpc/fams/orm/table"

	"gorm.io/gorm"
)

// BankAccountApplicationRepository 银行账户申请数据访问接口
type BankAccountApplicationRepository interface {
	// 创建银行账户申请
	Create(application *table.BankAccountApplication) error
	// 根据ID获取银行账户申请
	GetByID(id uint) (*table.BankAccountApplication, error)
	// 根据申请单号获取银行账户申请
	GetByApplicationNumber(applicationNumber string) (*table.BankAccountApplication, error)
	// 根据客户ID获取银行账户申请列表
	GetByCustomerID(customerID uint) ([]*table.BankAccountApplication, error)
	// 根据用户ID获取银行账户申请列表
	GetByUserID(userID uint) ([]*table.BankAccountApplication, error)
	// 根据客户经理ID获取银行账户申请列表
	GetByManagerID(managerID uint) ([]*table.BankAccountApplication, error)
	// 更新银行账户申请
	Update(application *table.BankAccountApplication) error
	// 删除银行账户申请
	Delete(id uint) error
	// 获取银行账户申请列表（分页）
	List(page, pageSize int64, keyword string, managerID uint, status string, applicationStatus int32, customerID uint, userID uint, auditorID uint, applicationTimeStart, applicationTimeEnd int64, auditTimeStart, auditTimeEnd int64, auditResult string) ([]*table.BankAccountApplication, int64, error)
	// 检查申请单号是否存在
	ExistsByApplicationNumber(applicationNumber string) (bool, error)
	// 更新申请状态
	UpdateStatus(id uint, status string) error
	// 审核申请
	AuditApplication(id uint, auditorID uint, auditorName, auditNote, auditResult string) error
}

// bankAccountApplicationRepository 银行账户申请数据访问实现
type bankAccountApplicationRepository struct {
	db *gorm.DB
}

// NewBankAccountApplicationRepository 创建银行账户申请Repository
func NewBankAccountApplicationRepository(db *gorm.DB) BankAccountApplicationRepository {
	return &bankAccountApplicationRepository{db: db}
}

func (r *bankAccountApplicationRepository) Create(application *table.BankAccountApplication) error {
	return r.db.Create(application).Error
}

func (r *bankAccountApplicationRepository) GetByID(id uint) (*table.BankAccountApplication, error) {
	var application table.BankAccountApplication
	err := r.db.First(&application, id).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *bankAccountApplicationRepository) GetByApplicationNumber(applicationNumber string) (*table.BankAccountApplication, error) {
	var application table.BankAccountApplication
	err := r.db.Where("application_number = ?", applicationNumber).First(&application).Error
	if err != nil {
		return nil, err
	}
	return &application, nil
}

func (r *bankAccountApplicationRepository) GetByCustomerID(customerID uint) ([]*table.BankAccountApplication, error) {
	var applications []*table.BankAccountApplication
	err := r.db.Where("customer_id = ?", customerID).Find(&applications).Error
	return applications, err
}

func (r *bankAccountApplicationRepository) GetByUserID(userID uint) ([]*table.BankAccountApplication, error) {
	var applications []*table.BankAccountApplication
	err := r.db.Where("user_id = ?", userID).Find(&applications).Error
	return applications, err
}

func (r *bankAccountApplicationRepository) GetByManagerID(managerID uint) ([]*table.BankAccountApplication, error) {
	var applications []*table.BankAccountApplication
	err := r.db.Where("manager_id = ?", managerID).Find(&applications).Error
	return applications, err
}

func (r *bankAccountApplicationRepository) Update(application *table.BankAccountApplication) error {
	return r.db.Save(application).Error
}

func (r *bankAccountApplicationRepository) Delete(id uint) error {
	return r.db.Delete(&table.BankAccountApplication{}, id).Error
}

func (r *bankAccountApplicationRepository) List(page, pageSize int64, keyword string, managerID uint, status string, applicationStatus int32, customerID uint, userID uint, auditorID uint, applicationTimeStart, applicationTimeEnd int64, auditTimeStart, auditTimeEnd int64, auditResult string) ([]*table.BankAccountApplication, int64, error) {
	var applications []*table.BankAccountApplication
	var total int64

	// 构建查询
	query := r.db.Model(&table.BankAccountApplication{})

	// 关键词搜索（申请单号、客户姓名、客户邮箱、审核人姓名）
	if keyword != "" {
		query = query.Where("application_number LIKE ? OR customer_name LIKE ? OR customer_email LIKE ? OR auditor_name LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 客户经理ID过滤
	if managerID > 0 {
		query = query.Where("manager_id = ?", managerID)
	}

	// 客户ID过滤
	if customerID > 0 {
		query = query.Where("customer_id = ?", customerID)
	}

	// 用户ID过滤
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}

	// 审核人ID过滤
	if auditorID > 0 {
		query = query.Where("auditor_id = ?", auditorID)
	}

	// 申请状态过滤（字符串状态）
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 申请状态过滤（数字状态，兼容旧版本）
	if applicationStatus >= 0 {
		var statusStr string
		switch applicationStatus {
		case 0:
			statusStr = "Submitted" // 待审核
		case 1:
			statusStr = "Approved" // 通过
		case 2:
			statusStr = "Rejected" // 拒绝
		}
		if statusStr != "" {
			query = query.Where("status = ?", statusStr)
		}
	}

	// 审核结果过滤
	if auditResult != "" {
		query = query.Where("audit_result = ?", auditResult)
	}

	// 申请时间范围筛选
	if applicationTimeStart > 0 {
		query = query.Where("application_time >= ?", applicationTimeStart)
	}
	if applicationTimeEnd > 0 {
		query = query.Where("application_time <= ?", applicationTimeEnd)
	}

	// 审核时间范围筛选
	if auditTimeStart > 0 {
		query = query.Where("audit_time >= ?", auditTimeStart)
	}
	if auditTimeEnd > 0 {
		query = query.Where("audit_time <= ?", auditTimeEnd)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("application_time DESC, id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&applications).Error; err != nil {
		return nil, 0, err
	}

	return applications, total, nil
}

func (r *bankAccountApplicationRepository) ExistsByApplicationNumber(applicationNumber string) (bool, error) {
	var count int64
	err := r.db.Model(&table.BankAccountApplication{}).Where("application_number = ?", applicationNumber).Count(&count).Error
	return count > 0, err
}

func (r *bankAccountApplicationRepository) UpdateStatus(id uint, status string) error {
	return r.db.Model(&table.BankAccountApplication{}).Where("id = ?", id).Update("status", status).Error
}

func (r *bankAccountApplicationRepository) AuditApplication(id uint, auditorID uint, auditorName, auditNote, auditResult string) error {
	return r.db.Model(&table.BankAccountApplication{}).Where("id = ?", id).Updates(map[string]interface{}{
		"auditor_id":   auditorID,
		"auditor_name": auditorName,
		"audit_note":   auditNote,
		"audit_result": auditResult,
		"audit_time":   "UNIX_TIMESTAMP()",
	}).Error
}
