package repository

import (
	"akatm/rpc/iam/orm/table"

	"gorm.io/gorm"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	// 创建用户
	Create(user *table.User) error
	// 根据ID获取用户
	GetByID(id uint) (*table.User, error)
	// 根据邮箱获取用户
	GetByEmail(email string) (*table.User, error)
	// 根据邀请码获取用户
	GetByInviteCode(inviteCode string) (*table.User, error)
	// 更新用户
	Update(user *table.User) error
	// 删除用户
	Delete(id uint) error
	// 获取用户列表（分页）
	List(page, pageSize int64, keyword string, isManager *bool) ([]*table.User, int64, error)
	// 检查邮箱是否存在
	ExistsByEmail(email string) (bool, error)
	// 检查邀请码是否存在
	ExistsByInviteCode(inviteCode string) (bool, error)
}

// userRepository 用户数据访问实现
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户Repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *table.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*table.User, error) {
	var user table.User
	err := r.db.Preload("Profile").Preload("Emails").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*table.User, error) {
	var user table.User
	err := r.db.Where("email = ?", email).Preload("Profile").Preload("Emails").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByInviteCode(inviteCode string) (*table.User, error) {
	var user table.User
	err := r.db.Where("invite_code = ?", inviteCode).Preload("Profile").Preload("Emails").First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *table.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&table.User{}, id).Error
}

func (r *userRepository) List(page, pageSize int64, keyword string, isManager *bool) ([]*table.User, int64, error) {
	var users []*table.User
	var total int64

	// 构建查询
	query := r.db.Model(&table.User{})

	// 关键词搜索
	if keyword != "" {
		query = query.Where("email LIKE ? OR invite_code LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 是否为客户经理过滤
	if isManager != nil {
		query = query.Where("is_manager = ?", *isManager)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Preload("Profile").Preload("Emails").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&table.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *userRepository) ExistsByInviteCode(inviteCode string) (bool, error) {
	var count int64
	err := r.db.Model(&table.User{}).Where("invite_code = ?", inviteCode).Count(&count).Error
	return count > 0, err
}
