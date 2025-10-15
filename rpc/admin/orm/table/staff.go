package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Name   string `gorm:"size:50;unique"`
	UserID uint   `gorm:"index"`
}

func init() {
	orm.RegisterTables(Staff{})
}
