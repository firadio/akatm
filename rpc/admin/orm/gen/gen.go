package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"akatm/rpc/admin/orm"
	_ "akatm/rpc/admin/orm/table"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func loadEnv(env string) map[string]string {
	fname := ".env"
	if env != "" {
		fname = ".env." + env
	}
	// resolve relative to this file directory
	base := filepath.Dir(os.Args[0])
	// when go run, working dir is this folder; fall back to current dir
	path := fname
	if _, err := os.Stat(path); err != nil {
		path = filepath.Join(base, fname)
	}
	m := map[string]string{}
	f, err := os.Open(path)
	if err != nil {
		return m
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) == 2 {
			m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}
	return m
}

func main() {
	env := os.Getenv("ENV") // dev/pro/test，默认读取 .env
	cfg := loadEnv(env)
	dsn := cfg["DSN"]
	if dsn == "" {
		panic("missing DSN in .env files")
	}
	tablePrefix := cfg["TABLE_PREFIX"]
	if tablePrefix == "" {
		tablePrefix = "admin_"
	}
	fmt.Println("[gen] using env=", env, "prefix=", tablePrefix)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 所有表加前缀
			SingularTable: false,       // 表名是否单数
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
