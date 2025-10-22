package table

import (
	"akatm/rpc/admin/orm"

	"gorm.io/gorm"
)

// User 员工
type SysUser struct {
	gorm.Model
	Name     string `gorm:"size:50;index"`
	Email    string `gorm:"size:50;uniqueIndex"`
	Password string `gorm:"size:255"`                     // 登录密码哈希
	Status   int8   `gorm:"type:tinyint;default:1;index"` // 1启用 0禁用
	// 关联的角色
	Roles []SysUserRole `gorm:"foreignKey:UserId"`
}

// Role 角色
type SysRole struct {
	gorm.Model
	Name        string `gorm:"size:50;uniqueIndex"`          // 角色名称
	Code        string `gorm:"size:50;uniqueIndex"`          // 角色代码
	Description string `gorm:"size:255"`                     // 角色描述
	Status      int8   `gorm:"type:tinyint;default:1;index"` // 1启用 0禁用
	// 关联的菜单
	Menus []SysRoleMenu `gorm:"foreignKey:RoleId"`
	// 关联的员工
	Users []SysUserRole `gorm:"foreignKey:RoleId"`
}

// Menu 菜单
type SysMenu struct {
	gorm.Model
	Name      string `gorm:"size:50"`                      // 菜单名称
	Path      string `gorm:"size:100"`                     // 菜单路径
	Icon      string `gorm:"size:50"`                      // 菜单图标
	Component string `gorm:"size:100"`                     // 组件路径
	Role      string `gorm:"size:50"`                      // 角色权限
	Label     string `gorm:"size:50"`                      // 显示标签
	Alias     string `gorm:"size:100"`                     // 别名路径
	Type      int8   `gorm:"type:tinyint;default:1;index"` // 类型：1菜单 2按钮
	ParentId  uint   `gorm:"index"`                        // 父菜单ID
	Sort      int    `gorm:"default:0"`                    // 排序
	Status    int8   `gorm:"type:tinyint;default:1;index"` // 1启用 0禁用
	// 子菜单
	Children []SysMenu `gorm:"foreignKey:ParentId"`
	// 关联的角色
	Roles []SysRoleMenu `gorm:"foreignKey:MenuId"`
}

// Permission 权限点
type SysPermission struct {
	gorm.Model
	Name        string `gorm:"size:50;uniqueIndex"`          // 权限名称
	Code        string `gorm:"size:50;uniqueIndex"`          // 权限代码（如 customer:read）
	Resource    string `gorm:"size:50"`                      // 资源（如 customer）
	Action      string `gorm:"size:50"`                      // 动作（如 read）
	Description string `gorm:"size:255"`                     // 权限描述
	Status      int8   `gorm:"type:tinyint;default:1;index"` // 1启用 0禁用
}

// UserRole 员工角色关联
type SysUserRole struct {
	gorm.Model
	UserId uint `gorm:"index"`
	RoleId uint `gorm:"index"`
	// 关联的员工
	User SysUser `gorm:"foreignKey:UserId"`
	// 关联的角色
	Role SysRole `gorm:"foreignKey:RoleId"`
}

// RoleMenu 角色菜单关联
type SysRoleMenu struct {
	gorm.Model
	RoleId uint `gorm:"index"`
	MenuId uint `gorm:"index"`
	// 关联的角色
	Role SysRole `gorm:"foreignKey:RoleId"`
	// 关联的菜单
	Menu SysMenu `gorm:"foreignKey:MenuId"`
}

// UserLog 员工操作日志
type SysUserLog struct {
	gorm.Model
	// 关联的员工ID
	UserId uint `gorm:"index"`
	// 操作类型（登录、登出、密码修改等）
	Action string `gorm:"size:50"`
	// 操作时间
	ActionTime int64 `gorm:"index"`
	// 操作IP地址
	OperatedIP string `gorm:"size:45"`
	// 操作描述
	Description string `gorm:"size:255"`
	// 备注
	Note string `gorm:"size:255"`
}

func init() {
	orm.RegisterTables(
		SysUser{},
		SysUserLog{},
		SysUserRole{},
		SysRole{},
		SysMenu{},
		SysPermission{},
		SysRoleMenu{},
	)
}
