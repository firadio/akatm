package table

import (
	"akatm/rpc/fams/orm"

	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	// 邀请码
	InviteCode string `gorm:"size:20;unique"`
	// 员工ID（关联Staff表）
	StaffId uint `gorm:"index;default:null"`
	// 关联的用户邮箱
	Emails []UserEmail `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户会话
	Sessions []UserSession `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户凭证
	Credentials []UserCredential `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户资料
	Profile UserProfile `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

// UserCredential 用户凭证
type UserCredential struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 凭证类型（如登录密码、资金密码、OAuth等）
	CredentialType string `gorm:"size:50"`
	// 凭证值（如哈希后的密码、OAuth令牌等）
	CredentialValue string `gorm:"size:255"`
}

// UserCredentialLog 用户凭证操作日志
type UserCredentialLog struct {
	gorm.Model
	// 关联的用户凭证ID
	UserCredentialId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 凭证类型（如登录密码、资金密码、OAuth等）
	CredentialType string `gorm:"size:50"`
	// 操作类型（创建、更新、删除）
	Action string `gorm:"size:20"`
	// 操作时间
	ActionTime int64 `gorm:"index"`
	// 旧凭证值（如哈希后的密码、OAuth令牌等）
	OldCredentialValue string `gorm:"size:255"`
	// 新凭证值（如哈希后的密码、OAuth令牌等）
	NewCredentialValue string `gorm:"size:255"`
	// 操作IP地址
	OperatedIP string `gorm:"size:45"`
	// 操作人员类型（用户、管理员等）
	OperatedByType string `gorm:"size:20"`
	// 操作人员ID
	OperatedByStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatedByStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserProfile 用户资料
type UserProfile struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 昵称
	Nickname string `gorm:"size:50"`
	// 头像URL
	AvatarURL string `gorm:"size:255"`
	// 其他扩展信息（JSON格式）
	ExtraInfo string `gorm:"type:text"`
}

// UserEmail 用户邮箱
type UserEmail struct {
	gorm.Model
	// 邮箱地址
	Email string `gorm:"size:50;unique"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
}

// UserEmailLog 用户邮箱绑定删除日志
type UserEmailLog struct {
	gorm.Model
	// 关联的用户邮箱ID
	UserEmailId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 操作类型（绑定、删除）
	Action string `gorm:"size:20"`
	// 操作时间
	ActionTime int64 `gorm:"index"`
	// 操作IP地址
	OperatedIP string `gorm:"size:45"`
	// 操作人员类型（用户、管理员等）
	OperatedByType string `gorm:"size:20"`
	// 操作人员ID
	OperatedByStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatedByStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserSession 用户会话
type UserSession struct {
	gorm.Model
	// Token 用于身份验证的令牌
	Token string `gorm:"size:64;uniqueIndex"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 过期时间
	ExpiresAt int64 `gorm:"index"`
}

// UserSessionLog 用户会话操作日志
type UserSessionLog struct {
	gorm.Model
	// 关联的用户会话ID
	UserSessionId uint `gorm:"index"`
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 操作类型（创建、更新、删除）
	Action string `gorm:"size:20"`
	// 操作时间
	ActionTime int64 `gorm:"index"`
	// 旧过期时间
	OldExpiresAt int64 `gorm:"index"`
	// 新过期时间
	NewExpiresAt int64 `gorm:"index"`
	// 操作IP地址
	OperatedIP string `gorm:"size:45"`
	// 操作人员类型（用户、管理员等）
	OperatedByType string `gorm:"size:20"`
	// 操作人员ID
	OperatedByStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatedByStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

func init() {
	orm.RegisterTables(User{}, UserEmail{})
}
