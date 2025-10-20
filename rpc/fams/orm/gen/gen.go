package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"akatm/rpc/fams/orm"
	_ "akatm/rpc/fams/orm/table"

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
	base := filepath.Dir(os.Args[0])
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
	env := os.Getenv("ENV")
	cfg := loadEnv(env)
	dsn := cfg["DSN"]
	if dsn == "" {
		// 使用默认DSN，仅用于生成代码
		dsn = "root:password@tcp(localhost:3306)/akatm_fams?charset=utf8mb4&parseTime=True&loc=Local"
		fmt.Println("[gen] using default DSN for code generation only")
	}
	tablePrefix := cfg["TABLE_PREFIX"]
	if tablePrefix == "" {
		tablePrefix = "fams_"
	}
	fmt.Println("[gen] using env=", env, "prefix=", tablePrefix)

	// 尝试连接数据库，如果失败则跳过
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 所有表加前缀
			SingularTable: false,       // 表名是否单数
		},
	})
	if err != nil {
		fmt.Printf("[gen] database connection failed: %v, skipping migration\n", err)
		// 创建一个内存数据库用于生成代码
		db, err = gorm.Open(mysql.Open("root:password@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   tablePrefix,
				SingularTable: false,
			},
		})
		if err != nil {
			fmt.Printf("[gen] fallback database connection also failed: %v\n", err)
			return
		}
	} else {
		// 自动迁移
		db.AutoMigrate(orm.AllTables...)
	}

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
