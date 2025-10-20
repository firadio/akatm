package table

import (
	"akatm/rpc/fams/orm"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// BankCustomer 银行客户
type BankCustomer struct {
	gorm.Model
	// 客户姓名
	Name string `gorm:"size:100"`
	// 手机号码
	Phone string `gorm:"size:20;uniqueIndex"`
	// 电子邮箱
	Email string `gorm:"size:50;uniqueIndex"`
	// 地址
	Address string `gorm:"size:255"`
	// 证件号（个人）
	IdNumber string `gorm:"size:64"`
	// 营业执照（企业）
	BusinessLicense string `gorm:"size:128"`
	// 客户类型（individual-个人, enterprise-企业）
	CustomerType string `gorm:"size:20;index"`
	// KYC状态（pending-待审核, verified-已认证, rejected-已拒绝）
	KycStatus string `gorm:"size:20;index"`
	// KYC认证时间
	KycVerifiedAt int64 `gorm:"index"`
	// KYC认证材料（JSON格式）
	KycDocuments string `gorm:"type:text"`
	// 备注
	Note string `gorm:"size:255"`
	// 关联的用户ID（系统用户）
	UserId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`

	// 统计字段
	BankAccountCount int32           `gorm:"default:0"`                    // 银行账户数量
	WalletCount      int32           `gorm:"default:0"`                    // 钱包数量
	TotalBalance     decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 总余额

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
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 账户状态 1-正常 0-禁用
	Status int8 `gorm:"type:tinyint;default:1;index"`

	// 余额相关
	Balance          decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 账户余额
	AvailableBalance decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 可用余额
	FrozenBalance    decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 冻结余额

	// 交易统计
	TransactionCount      int64           `gorm:"default:0"`                    // 交易笔数
	TotalDepositAmount    decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 累计存款金额
	TotalWithdrawalAmount decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 累计提现金额
	TotalFeeAmount        decimal.Decimal `gorm:"type:decimal(20,8);default:0"` // 累计手续费

	// 限额设置
	DailyTransactionLimit   decimal.Decimal `gorm:"type:decimal(20,8)"` // 日交易限额
	MonthlyTransactionLimit decimal.Decimal `gorm:"type:decimal(20,8)"` // 月交易限额

	// 时间字段
	LastTransactionAt  int64 `gorm:"index"` // 最后交易时间
	AccountOpeningDate int64 `gorm:"index"` // 开户日期
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
	DepositAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 存款时间
	DepositTime int64 `gorm:"index"`
	// 银行入账状态
	DepositStatus string `gorm:"size:50"`
	// 审核状态（待审核、已审核）
	AuditStatus int8 `gorm:"type:tinyint;default:0"`

	// 换汇相关
	ExchangeAmount decimal.Decimal `gorm:"type:decimal(20,8)"` // 换汇后金额
	ExchangeFee    decimal.Decimal `gorm:"type:decimal(20,8)"` // 换汇手续费
	ExchangeRate   decimal.Decimal `gorm:"type:decimal(20,8)"` // 汇率

	// 审核信息
	AuditorId   uint   `gorm:"index"`    // 审核人ID
	AuditorName string `gorm:"size:50"`  // 审核人姓名
	AuditTime   int64  `gorm:"index"`    // 审核时间
	AuditNote   string `gorm:"size:255"` // 审核意见
	AuditResult string `gorm:"size:20"`  // 审核结果

	// 状态字段
	Status string `gorm:"size:20;index"` // 状态
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

// BankAccountApplication 银行账户开户申请
type BankAccountApplication struct {
	gorm.Model
	// 申请单号
	ApplicationNumber string `gorm:"size:30;uniqueIndex"`
	// 关联的客户ID
	CustomerId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 申请状态（Draft、Submitted、UnderReview、Approved、Rejected）
	Status string `gorm:"size:20;index"`
	// 申请材料（JSON格式）
	Materials string `gorm:"type:text"`
	// 申请时间
	ApplicationTime int64 `gorm:"index"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人姓名
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 审核意见
	AuditNote string `gorm:"size:255"`
	// 审核结果
	AuditResult string `gorm:"size:20"`
	// 备注
	Note string `gorm:"size:255"`
}

// BankWebhookRecord 银行回调记录
type BankWebhookRecord struct {
	gorm.Model
	// 幂等键（交易流水号+渠道）
	IdempotencyKey string `gorm:"size:64;uniqueIndex"`
	// 银行渠道
	Channel string `gorm:"size:50"`
	// 签名
	Signature string `gorm:"size:255"`
	// Payload 哈希
	PayloadHash string `gorm:"size:64"`
	// 原始 Payload
	RawPayload string `gorm:"type:text"`
	// 处理状态（Pending、Success、Failed）
	ProcessStatus string `gorm:"size:20;index"`
	// 处理时间
	ProcessTime int64 `gorm:"index"`
	// 重试次数
	RetryCount int `gorm:"default:0"`
	// 最后重试时间
	LastRetryTime int64 `gorm:"index"`
	// 错误信息
	ErrorMessage string `gorm:"size:255"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserWalletDeposit 用户钱包存款记录
type UserWalletDeposit struct {
	gorm.Model
	// 存款单号
	DepositNumber string `gorm:"size:30;uniqueIndex"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 关联的银行账户ID
	BankAccountId uint `gorm:"index"`
	// 关联的银行存款记录ID
	BankDepositId uint `gorm:"index"`
	// 存款金额
	DepositAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 换汇后金额
	ExchangeAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 换汇手续费
	ExchangeFee decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 存款时间
	DepositTime int64 `gorm:"index"`
	// 审核状态（Pending、Approved、Rejected）
	AuditStatus string `gorm:"size:20;index"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人姓名
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 审核意见
	AuditNote string `gorm:"size:255"`
	// 审核结果
	AuditResult string `gorm:"size:20"`
	// 备注
	Note string `gorm:"size:255"`
}

// BankWithdrawal 银行账户提现记录
type BankWithdrawal struct {
	gorm.Model
	// 提现单号
	WithdrawNumber string `gorm:"size:30;uniqueIndex"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的客户ID
	CustomerId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 关联的银行账户ID
	BankAccountId uint `gorm:"index"`
	// 提现金额（USDT）
	WithdrawAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 提现手续费
	WithdrawFee decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 实际到账金额
	ActualAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 目标虚拟钱包地址
	CryptoAddress string `gorm:"size:255"`
	// 网络类型（如：ERC20、TRC20、BEP20等）
	Network string `gorm:"size:10"`
	// 地址标签/备注
	Tag string `gorm:"size:100"`
	// 提现状态（Requested-已申请、UnderReview-审核中、Approved-已通过、Rejected-已拒绝、Settled-已结算）
	Status string `gorm:"size:20;index"`
	// 提现时间
	WithdrawTime int64 `gorm:"index"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人姓名
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 审核意见
	AuditNote string `gorm:"size:255"`
	// 审核结果
	AuditResult string `gorm:"size:20"`
	// 备注
	Note string `gorm:"size:255"`
}

// FundDetail 资金明细记录
type FundDetail struct {
	gorm.Model
	// 交易单号
	TransactionNumber string `gorm:"size:30;uniqueIndex"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 用户类型（super_agent、agent、manager、customer）
	UserType string `gorm:"size:20;index"`
	// 父级用户ID
	ParentUserId uint `gorm:"index"`
	// 交易类型（deposit、withdrawal、exchange、fee）
	TransactionType string `gorm:"size:20;index"`
	// 交易金额
	Amount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 币种
	Currency string `gorm:"size:10;index"`
	// 手续费
	Fee decimal.Decimal `gorm:"type:decimal(20,8);default:0"`
	// 实际到账金额
	ActualAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 状态（pending、completed、failed、cancelled）
	Status string `gorm:"size:20;index"`
	// 描述
	Description string `gorm:"size:255"`
	// 关联的存款记录ID（如果是存款交易）
	DepositId uint `gorm:"index;default:null"`
	// 关联的提现记录ID（如果是提现交易）
	WithdrawalId uint `gorm:"index;default:null"`
	// 交易时间
	TransactionTime int64 `gorm:"index"`
	// 备注
	Note string `gorm:"size:255"`
}

func init() {
	orm.RegisterTables(
		BankCustomer{},
		BankAccount{},
		BankDeposit{},
		BankDepositAudit{},
		BankAccountApplication{},
		BankWebhookRecord{},
		UserWalletDeposit{},
		BankWithdrawal{},
		FundDetail{},
	)
}
