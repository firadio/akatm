package orm

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 用来收集所有需要建表和生成 CRUD 的 struct
var AllTables []any

// 注册函数
func RegisterTables(tables ...any) {
	AllTables = append(AllTables, tables...)
}

// NewDB 创建数据库连接
func NewDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移所有表
	if err := db.AutoMigrate(AllTables...); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	fmt.Printf("Database connected successfully, migrated %d tables\n", len(AllTables))
	return db
}
