package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

// UserSetting 用户设置
type UserSetting struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 换汇手续费率
	ExchangeFeeRate float64 `gorm:"type:decimal(5,4);default:0.0600"`
	// 提现手续费
	WithdrawFee float64 `gorm:"type:decimal(10,2);default:2.00"`
}

// UserWallet 用户钱包
type UserWallet struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 资产代号 (如 USDT, BTC)
	Symbol string `gorm:"size:10"`
	// 可用余额
	AvailableBalance float64 `gorm:"type:decimal(20,8);default:0.00000000"`
	// 冻结余额
	FrozenAmount float64 `gorm:"type:decimal(20,8);default:0.00000000"`
}

// UserWalletLog 用户钱包变动日志
type UserWalletLog struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 币种 (如 USD, EUR)
	Currency string `gorm:"size:10"`
	// 资产代号 (如 BTC, ETH)
	Symbol string `gorm:"size:10"`
	// 账变类型（如 1:充值, 2:提现, 3:交易等）
	TranscationType string `gorm:"size:50"`
	// 账变名称（如 存款成功、发起提现、提现审核失败、提现审核成功）
	TranscationName string `gorm:"size:50"`
	// 账变金额（不能为负数）
	TranscationAmount float64 `gorm:"type:decimal(20,8)"`
	// 变动前可用余额
	BeforeAvailableBalance float64 `gorm:"type:decimal(20,8)"`
	// 变动后可用余额
	AfterAvailableBalance float64 `gorm:"type:decimal(20,8)"`
	// 变动前冻结余额
	BeforeFrozenAmount float64 `gorm:"type:decimal(20,8)"`
	// 变动后冻结余额
	AfterFrozenAmount float64 `gorm:"type:decimal(20,8)"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserBalance 用户余额
type UserBalance struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 可用余额
	AvailableBalance float64 `gorm:"type:decimal(20,2);default:0.00"`
	// 冻结余额
	FrozenAmount float64 `gorm:"type:decimal(20,2);default:0.00"`
}

// 用户余额变动日志
type UserBalanceLog struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 变动类型（如充值、提现、消费等）
	ChangeType string `gorm:"size:50"`
	// 变动金额
	ChangeAmount float64 `gorm:"type:decimal(20,2)"`
	// 变动前余额
	BeforeBalance float64 `gorm:"type:decimal(20,2)"`
	// 变动后余额
	AfterBalance float64 `gorm:"type:decimal(20,2)"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserWithdrawal 用户提现记录
type UserWithdrawal struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 提现金额
	WithdrawAmount float64 `gorm:"type:decimal(20,2)"`
	// 提现状态（待处理、已处理、拒绝等）
	Status string `gorm:"size:50"`
	// 提现时间
	WithdrawTime int64 `gorm:"index"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserWithdrawalAddress 用户提现地址
type UserWithdrawalAddress struct {
	gorm.Model
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
	// 是否默认地址
	IsDefault bool `gorm:"default:false"`
	// 审核状态（待审核、已审核、拒绝）
	AuditStatus int8 `gorm:"type:tinyint;default:0"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserWithdrawalAddressAudit 用户提现地址审核记录
type UserWithdrawalAddressAudit struct {
	gorm.Model
	// 关联的提现地址ID
	WithdrawalAddressId uint `gorm:"index"`
	// 审核人ID
	AuditorId uint `gorm:"index"`
	// 审核人名字
	AuditorName string `gorm:"size:50"`
	// 审核时间
	AuditTime int64 `gorm:"index"`
	// 审核备注
	AuditNote string `gorm:"size:255"`
	// 审核结果（通过、拒绝）
	AuditResult string `gorm:"size:20"`
}

// UserWithdrawalAddressLog 用户提现地址变动日志
type UserWithdrawalAddressLog struct {
	gorm.Model
	// 关联的提现地址ID
	WithdrawalAddressId uint `gorm:"index"`
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
	OperatedIP string `gorm:"size:45"`
	// 操作人员类型（用户、管理员等）
	OperatedByType string `gorm:"size:20"`
	// 操作人员ID
	OperatedByStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatedByStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// Init 注册表
func init() {
	orm.RegisterTables(UserSetting{})
}
