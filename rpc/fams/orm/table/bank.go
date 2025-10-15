package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

// BankCustomer 银行客户
type BankCustomer struct {
	gorm.Model
	// 名字
	FirstName string `gorm:"size:50"`
	// 姓氏
	LastName string `gorm:"size:64"`
	// 手机号码
	Phone string `gorm:"size:20;uniqueIndex"`
	// 电子邮箱
	Email string `gorm:"size:50;uniqueIndex"`
	// 地址
	Address string `gorm:"size:255"`
	// 备注
	Note string `gorm:"size:255"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的银行账户
	BankAccounts []BankAccount `gorm:"foreignKey:CustomerId"`
}

// BankAccount 银行账户
type BankAccount struct {
	gorm.Model
	// 银行账号
	AccountNumber string `gorm:"size:20;unique"`
	// 币种
	Currency string `gorm:"size:10"`
	// 关联的客户ID
	CustomerId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
}

// BankDeposit 银行存款记录
type BankDeposit struct {
	gorm.Model
	// 存款单号
	DepositNumber string `gorm:"size:30;unique"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的银行账户ID
	BankAccountId uint `gorm:"index"`
	// 存款金额
	DepositAmount float64 `gorm:"type:decimal(20,2)"`
	// 存款时间
	DepositTime int64 `gorm:"index"`
	// 银行入账状态
	DepositStatus string `gorm:"size:50"`
	// 审核状态（待审核、已审核）
	AuditStatus int8 `gorm:"type:tinyint;default:0"`
}

// BankDepositAudit 银行存款审核记录
type BankDepositAudit struct {
	gorm.Model
	// 关联的存款记录ID
	BankDepositId uint `gorm:"index"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人名字
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 审核备注
	AuditNote string `gorm:"size:255"`
	// 审核结果（通过、拒绝）
	AuditResult string `gorm:"size:50"`
}

func init() {
	orm.RegisterTables(BankAccount{})
}
