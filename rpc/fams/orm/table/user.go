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

// UserBalance 用户余额
type UserBalance struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 可用余额
	AvailableBalance float64 `gorm:"type:decimal(20,2);default:0.00"`
	// 冻结余额
	HoldBalance float64 `gorm:"type:decimal(20,2);default:0.00"`
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

func init() {
	orm.RegisterTables(UserSetting{})
}
