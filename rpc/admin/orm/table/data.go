package table

import (
	"akatm/rpc/admin/orm"

	"gorm.io/gorm"
)

// Country 开户国家
type DataCountry struct {
	gorm.Model
	Code        string `gorm:"size:10;uniqueIndex"`          // 国家代码（如 CN, US）
	Name        string `gorm:"size:100;index"`               // 国家名称
	NameEn      string `gorm:"size:100;index"`               // 英文名称
	Currency    string `gorm:"size:10"`                      // 货币代码
	PhoneCode   string `gorm:"size:10"`                      // 电话区号
	Sort        int    `gorm:"default:0"`                    // 排序
	Status      int8   `gorm:"type:tinyint;default:1;index"` // 1启用 0禁用
	Description string `gorm:"size:255"`                     // 描述
}

func init() {
	orm.RegisterTables(
		DataCountry{},
	)
}
