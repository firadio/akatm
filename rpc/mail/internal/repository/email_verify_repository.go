package repository

import (
	"akatm/rpc/mail/orm/table"

	"gorm.io/gorm"
)

// EmailVerifyRepository 邮件验证数据访问接口
type EmailVerifyRepository interface {
	// 创建邮件验证记录
	Create(emailVerify *table.EmailVerify) error
	// 根据邮箱获取验证记录
	GetByEmail(email string) (*table.EmailVerify, error)
	// 根据邮箱和场景获取验证记录
	GetByEmailAndScene(email, scene string) (*table.EmailVerify, error)
	// 更新邮件验证记录
	Update(emailVerify *table.EmailVerify) error
	// 删除邮件验证记录
	Delete(email string) error
	// 删除过期的验证记录
	DeleteExpired() error
	// 检查验证码是否正确
	VerifyCode(email, scene, code string) (bool, error)
	// 检查是否在频率限制内
	CheckFrequencyLimit(email, scene string, limitSeconds int64) (bool, error)
}

// emailVerifyRepository 邮件验证数据访问实现
type emailVerifyRepository struct {
	db *gorm.DB
}

// NewEmailVerifyRepository 创建邮件验证Repository
func NewEmailVerifyRepository(db *gorm.DB) EmailVerifyRepository {
	return &emailVerifyRepository{db: db}
}

func (r *emailVerifyRepository) Create(emailVerify *table.EmailVerify) error {
	return r.db.Create(emailVerify).Error
}

func (r *emailVerifyRepository) GetByEmail(email string) (*table.EmailVerify, error) {
	var emailVerify table.EmailVerify
	err := r.db.Where("email = ?", email).First(&emailVerify).Error
	if err != nil {
		return nil, err
	}
	return &emailVerify, nil
}

func (r *emailVerifyRepository) GetByEmailAndScene(email, scene string) (*table.EmailVerify, error) {
	var emailVerify table.EmailVerify
	err := r.db.Where("email = ? AND scene = ?", email, scene).First(&emailVerify).Error
	if err != nil {
		return nil, err
	}
	return &emailVerify, nil
}

func (r *emailVerifyRepository) Update(emailVerify *table.EmailVerify) error {
	return r.db.Save(emailVerify).Error
}

func (r *emailVerifyRepository) Delete(email string) error {
	return r.db.Where("email = ?", email).Delete(&table.EmailVerify{}).Error
}

func (r *emailVerifyRepository) DeleteExpired() error {
	// 删除过期的验证记录
	return r.db.Where("expires_at < ?", "UNIX_TIMESTAMP()").Delete(&table.EmailVerify{}).Error
}

func (r *emailVerifyRepository) VerifyCode(email, scene, code string) (bool, error) {
	var count int64
	err := r.db.Model(&table.EmailVerify{}).
		Where("email = ? AND scene = ? AND code = ? AND expires_at > ?",
			email, scene, code, "UNIX_TIMESTAMP()").
		Count(&count).Error
	return count > 0, err
}

func (r *emailVerifyRepository) CheckFrequencyLimit(email, scene string, limitSeconds int64) (bool, error) {
	var count int64
	err := r.db.Model(&table.EmailVerify{}).
		Where("email = ? AND scene = ? AND created_at > ?",
			email, scene, "UNIX_TIMESTAMP() - ?", limitSeconds).
		Count(&count).Error
	return count == 0, err
}
