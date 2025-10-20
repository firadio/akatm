package table

import (
	"akatm/rpc/iam/orm"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// User 前台用户表
type User struct {
	gorm.Model
	// 用户类型：super_agent(总代)、agent(代理)、manager(客户经理)、customer(用户)
	UserType string `gorm:"size:20;index"`
	// 父级用户ID（建立层级关系）
	ParentId uint `gorm:"index"`
	// 父级用户树（存储完整的父级路径，如：1,2,3）
	ParentTree string `gorm:"size:255;index"`
	// 主邮箱（冗余字段，便于快速查询；与 UserEmail 保持一致）
	Email string `gorm:"size:50;index"`
	// 邀请码（注册时使用的邀请码）
	InviteCode string `gorm:"size:20;unique"`
	// 员工ID（关联Staff表，仅客户经理有此字段）
	StaffId uint `gorm:"index;default:null"`
	// 换汇手续费率（百分比，如5.00表示5%）
	ExchangeFeeRate decimal.Decimal `gorm:"type:decimal(5,4);default:0.0000"`
	// 提现手续费（固定金额，如5.00表示5美元）
	WithdrawFee decimal.Decimal `gorm:"type:decimal(10,2);default:0.00"`
	// 状态：1启用 0禁用
	Status int8 `gorm:"type:tinyint;default:1;index"`
	// 关联的用户邮箱
	Emails []UserEmail `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户会话
	Sessions []UserSession `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户凭证
	Credentials []UserCredential `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的用户资料
	Profile UserProfile `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 关联的授权国家
	AuthorizedCountries []UserCountryAuth `gorm:"foreignKey:UserId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// 子级用户（用于查询下级用户）
	Children []User `gorm:"foreignKey:ParentId"`
	// 父级用户（用于查询上级用户）
	Parent User `gorm:"foreignKey:ParentId"`
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

// UserInvite 用户邀请记录
type UserInvite struct {
	gorm.Model
	// 邀请码
	InviteCode string `gorm:"size:20;uniqueIndex"`
	// 生成邀请的用户ID（总代或代理）
	InviterUserId uint `gorm:"index"`
	// 生成邀请的用户类型
	InviterUserType string `gorm:"size:20"`
	// 生成邀请的用户邮箱
	InviterEmail string `gorm:"size:50"`
	// 目标用户类型（super_agent、agent、manager、customer）
	TargetUserType string `gorm:"size:20"`
	// 换汇手续费率（百分比，如5.00表示5%）
	ExchangeFeeRate decimal.Decimal `gorm:"type:decimal(5,4);default:0.0000"`
	// 提现手续费（固定金额，如5.00表示5美元）
	WithdrawFee decimal.Decimal `gorm:"type:decimal(10,2);default:0.00"`
	// 过期时间
	ExpiresAt int64 `gorm:"index"`
	// 是否已使用
	IsUsed bool `gorm:"default:false;index"`
	// 使用时间
	UsedAt int64 `gorm:"index"`
	// 使用人ID（注册的用户ID）
	UsedByUserId uint `gorm:"index"`
	// 使用人邮箱
	UsedByEmail string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserInviteLog 用户邀请操作日志
type UserInviteLog struct {
	gorm.Model
	// 关联的邀请记录ID
	UserInviteId uint `gorm:"index"`
	// 操作类型（生成、使用、过期）
	Action string `gorm:"size:20"`
	// 操作时间
	ActionTime int64 `gorm:"index"`
	// 操作IP地址
	OperatedIP string `gorm:"size:45"`
	// 操作人员类型（管理员、系统等）
	OperatedByType string `gorm:"size:20"`
	// 操作人员ID
	OperatedByStaffId uint `gorm:"index"`
	// 操作人员名字
	OperatedByStaffName string `gorm:"size:50"`
	// 备注
	Note string `gorm:"size:255"`
}

// UserCountryAuth 用户国家授权
type UserCountryAuth struct {
	gorm.Model
	// 关联的用户ID
	UserId uint `gorm:"index"`
	// 关联的国家ID
	CountryId uint `gorm:"index"`
	// 状态：1启用 0禁用
	Status int8 `gorm:"type:tinyint;default:1;index"`
	// 关联的国家信息
	Country Country `gorm:"foreignKey:CountryId"`
}

// Country 开户国家（与admin模块保持一致）
type Country struct {
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
		User{},
		UserCredential{},
		UserCredentialLog{},
		UserProfile{},
		UserEmail{},
		UserEmailLog{},
		UserSession{},
		UserSessionLog{},
		UserInvite{},
		UserInviteLog{},
		UserCountryAuth{},
		Country{},
	)
}
