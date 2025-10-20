package table

import (
	"akatm/rpc/iam/orm"

	"gorm.io/gorm"
)

type EmailVerify struct {
	gorm.Model
	// 邮箱地址
	Email string `gorm:"size:50;unique"`
	// 场景 (注册, 重置密码等)
	Scene string `gorm:"size:20;index"`
	// 验证码
	Code string `gorm:"size:6"`
	// 过期时间
	ExpiresAt int64 `gorm:"index"`
}

func init() {
	orm.RegisterTables(EmailVerify{})
}
