package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

type Customer struct {
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
}

func init() {
	orm.RegisterTables(Customer{})
}
