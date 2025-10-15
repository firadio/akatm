package orm

// 用来收集所有需要建表和生成 CRUD 的 struct
var AllTables []any

// 注册函数
func RegisterTables(tables ...any) {
	AllTables = append(AllTables, tables...)
}
