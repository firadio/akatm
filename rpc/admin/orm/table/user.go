package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string      `gorm:"size:50;unique"`
	Password string      `gorm:"size:64"`
	Emails   []UserEmail `gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

type UserEmail struct {
	gorm.Model
	Email  string `gorm:"size:50;unique"`
	UserID uint   `gorm:"index"`
}
type EmailVerify struct {
	gorm.Model
	Email string `gorm:"size:50;unique"`
	Code  string `gorm:"size:6"`
}

func init() {
	orm.RegisterTables(User{}, UserEmail{}, EmailVerify{})
}
