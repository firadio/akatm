package main

import (
	"akatm/rpc/fams/orm"
	_ "akatm/rpc/fams/orm/table"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// 连接数据库
	dsn := "akatm:akatm@tcp([240b:4001:278:8403:0:1acf:dae9:67eb]:3320)/akatm"
	dsn += "?charset=utf8mb4&parseTime=True&loc=UTC"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "fams_", // 所有表加前缀
			SingularTable: false,   // 表名是否单数
		},
	})
	if err != nil {
		panic(err)
	}

	// 自动迁移
	db.AutoMigrate(orm.AllTables...)

	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./query", // 输出路径
		ModelPkgPath:  "./model", // 模型文件路径
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.UseDB(db)

	// Generate the models
	var models []any
	for _, model := range orm.AllTables {
		stmt := &gorm.Statement{DB: db}
		stmt.Parse(model)
		models = append(models, g.GenerateModel(stmt.Schema.Table))
	}
	g.ApplyBasic(models...)

	// Execute the generator
	g.Execute()
}
