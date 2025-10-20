package table

import (
	"akatm/rpc/fams/orm"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// UserSetting 用户设置
type UserSetting struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 换汇手续费率
	ExchangeFeeRate decimal.Decimal `gorm:"type:decimal(5,4);default:0.0600"`
	// 提现手续费
	WithdrawFee decimal.Decimal `gorm:"type:decimal(10,2);default:2.00"`
}

// UserWallet 用户钱包
type UserWallet struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的客户ID（业务客户）
	CustomerId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 资产代号 (如 USDT, BTC)
	Symbol string `gorm:"index;size:10;default:'USDT'"`
	// 用户提现地址列表
	UserWalletAddresses []UserWalletAddress `gorm:"foreignKey:UserWalletId"`
	// 可用余额
	AvailableBalance decimal.Decimal `gorm:"type:decimal(20,8);default:0.00000000"`
	// 冻结余额
	FrozenAmount decimal.Decimal `gorm:"type:decimal(20,8);default:0.00000000"`
}

type TransactionType string

const (
	TransactionDeposit  TransactionType = "deposit"
	TransactionWithdraw TransactionType = "withdraw"
	TransactionTrade    TransactionType = "trade"
)

// UserWalletLedger 用户钱包账变记录
type UserWalletLedger struct {
	gorm.Model

	// User ID
	UserId uint `gorm:"index;not null"`

	// 用户钱包ID
	UserWalletId uint `gorm:"index;not null"`

	// 资产代号 (如 USDT, BTC)
	Symbol string `gorm:"index;size:10;default:'USDT'"`

	// Ledger type / business type (e.g. deposit, withdraw, trade)
	TransactionType TransactionType `gorm:"size:50;not null"`

	// Ledger name / display name (e.g. Deposit Success, Withdraw Initiated)
	TransactionName string `gorm:"size:100"`

	// Amount changed (always positive)
	Amount decimal.Decimal `gorm:"type:decimal(20,8);not null"`

	// Direction of change: +1 = increase, -1 = decrease
	ChangeDirection int8 `gorm:"not null;default:1"`

	// Before / after available balance
	BeforeAvailableBalance decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	AfterAvailableBalance  decimal.Decimal `gorm:"type:decimal(20,8);not null"`

	// Before / after frozen balance
	BeforeFrozenBalance decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	AfterFrozenBalance  decimal.Decimal `gorm:"type:decimal(20,8);not null"`

	// Reference transaction or order ID
	ReferenceId string `gorm:"size:64"`

	// Remarks
	Note string `gorm:"size:255"`
}

// UserWalletWithdrawal 用户钱包提现记录
type UserWalletWithdrawal struct {
	gorm.Model
	// 提现单号
	WithdrawNumber string `gorm:"size:30;uniqueIndex"`
	// 关联的用户钱包ID
	UserWalletId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的客户经理ID
	ManagerId uint `gorm:"index"`
	// 关联的钱包地址ID
	WalletAddressId uint `gorm:"index"`
	// 提现金额
	WithdrawAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 提现手续费
	WithdrawFee decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 实际到账金额
	ActualAmount decimal.Decimal `gorm:"type:decimal(20,8)"`
	// 提现状态（Requested、UnderReview、Approved、Rejected、Settled）
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

// UserWalletAddress 用户钱包地址
type UserWalletAddress struct {
	gorm.Model
	// 用户钱包ID
	UserWalletId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 地址类型（如 提现地址、收款地址）默认提现地址
	AddressType string `gorm:"size:50;default:'withdrawal'"`
	// 网络类型 (如 ERC20, TRC20)
	Network string `gorm:"size:10"`
	// 提现地址
	Address string `gorm:"size:255"`
	// 标签/备注
	Tag string `gorm:"size:100"`
	// 是否默认地址
	IsDefault bool `gorm:"default:false"`
	// 审核状态（待审核、已审核、拒绝）
	AuditStatus int8 `gorm:"type:tinyint;default:0"`
	// 审核时间
	AuditTime time.Time `gorm:"index"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserWalletAddressAudit 用户钱包地址审核记录
type UserWalletAddressAudit struct {
	gorm.Model
	// 关联的钱包地址ID
	UserWalletAddressId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 变动前的值
	OldValue string `gorm:"type:text"`
	// 变动后的值
	NewValue string `gorm:"type:text"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人名字
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime time.Time `gorm:"index"`
	// 审核备注
	AuditNote string `gorm:"size:255"`
	// 审核结果（通过、拒绝）
	AuditResult string `gorm:"size:20"`
}

// UserWalletAddressLog 用户提现地址变动日志
type UserWalletAddressLog struct {
	gorm.Model
	// 关联的钱包地址ID
	UserWalletAddressId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 币种 (如 USDT, BTC)
	Symbol string `gorm:"size:10"`
	// 网络类型 (如 ERC20, TRC20)
	Network string `gorm:"size:10"`
	// 提现地址
	Address string `gorm:"size:255"`
	// 标签/备注
	Tag string `gorm:"size:100"`
	// 变动类型（添加、删除）
	ChangeType string `gorm:"size:50"`
	// 操作IP地址
	OperatorIP string `gorm:"size:45"`
	// 操作人员类型（用户、管理员等）
	OperatorType string `gorm:"size:20"`
	// 操作人员ID
	OperatorStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatorStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// Init 注册表
func init() {
	orm.RegisterTables(
		UserSetting{},
		UserWallet{},
		UserWalletLedger{},
		UserWalletWithdrawal{},
		UserWalletAddress{},
		UserWalletAddressAudit{},
		UserWalletAddressLog{},
	)
}
