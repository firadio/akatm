package repository

import (
	"akatm/rpc/iam/orm/table"

	"gorm.io/gorm"
)

// UserInviteRepository 用户邀请数据访问接口
type UserInviteRepository interface {
	// 创建邀请
	Create(invite *table.UserInvite) error
	// 根据ID获取邀请
	GetByID(id uint) (*table.UserInvite, error)
	// 根据邀请码获取邀请
	GetByInviteCode(inviteCode string) (*table.UserInvite, error)
	// 更新邀请
	Update(invite *table.UserInvite) error
	// 删除邀请
	Delete(id uint) error
	// 获取邀请列表（分页）
	List(page, pageSize int64, staffID uint, isUsed *bool) ([]*table.UserInvite, int64, error)
	// 检查邀请码是否存在
	ExistsByInviteCode(inviteCode string) (bool, error)
	// 标记邀请为已使用
	MarkAsUsed(inviteCode string, userID uint, userEmail string) error
}

// userInviteRepository 用户邀请数据访问实现
type userInviteRepository struct {
	db *gorm.DB
}

// NewUserInviteRepository 创建用户邀请Repository
func NewUserInviteRepository(db *gorm.DB) UserInviteRepository {
	return &userInviteRepository{db: db}
}

func (r *userInviteRepository) Create(invite *table.UserInvite) error {
	return r.db.Create(invite).Error
}

func (r *userInviteRepository) GetByID(id uint) (*table.UserInvite, error) {
	var invite table.UserInvite
	err := r.db.First(&invite, id).Error
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func (r *userInviteRepository) GetByInviteCode(inviteCode string) (*table.UserInvite, error) {
	var invite table.UserInvite
	err := r.db.Where("invite_code = ?", inviteCode).First(&invite).Error
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func (r *userInviteRepository) Update(invite *table.UserInvite) error {
	return r.db.Save(invite).Error
}

func (r *userInviteRepository) Delete(id uint) error {
	return r.db.Delete(&table.UserInvite{}, id).Error
}

func (r *userInviteRepository) List(page, pageSize int64, staffID uint, isUsed *bool) ([]*table.UserInvite, int64, error) {
	var invites []*table.UserInvite
	var total int64

	// 构建查询
	query := r.db.Model(&table.UserInvite{})

	// 员工ID过滤
	if staffID > 0 {
		query = query.Where("staff_id = ?", staffID)
	}

	// 是否已使用过滤
	if isUsed != nil {
		query = query.Where("is_used = ?", *isUsed)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(int(offset)).Limit(int(pageSize)).Find(&invites).Error; err != nil {
		return nil, 0, err
	}

	return invites, total, nil
}

func (r *userInviteRepository) ExistsByInviteCode(inviteCode string) (bool, error) {
	var count int64
	err := r.db.Model(&table.UserInvite{}).Where("invite_code = ?", inviteCode).Count(&count).Error
	return count > 0, err
}

func (r *userInviteRepository) MarkAsUsed(inviteCode string, userID uint, userEmail string) error {
	return r.db.Model(&table.UserInvite{}).
		Where("invite_code = ?", inviteCode).
		Updates(map[string]interface{}{
			"is_used":         true,
			"used_at":         "UNIX_TIMESTAMP()",
			"used_by_user_id": userID,
			"used_by_email":   userEmail,
		}).Error
}
