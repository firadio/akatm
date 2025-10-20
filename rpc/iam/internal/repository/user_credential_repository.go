package repository

import (
	"akatm/rpc/iam/orm/table"

	"gorm.io/gorm"
)

// UserCredentialRepository 用户凭证数据访问接口
type UserCredentialRepository interface {
	// 创建凭证
	Create(credential *table.UserCredential) error
	// 根据ID获取凭证
	GetByID(id uint) (*table.UserCredential, error)
	// 根据用户ID和类型获取凭证
	GetByUserIDAndType(userID uint, credentialType string) (*table.UserCredential, error)
	// 更新凭证
	Update(credential *table.UserCredential) error
	// 删除凭证
	Delete(id uint) error
	// 根据用户ID删除所有凭证
	DeleteByUserID(userID uint) error
	// 检查用户是否有特定类型的凭证
	ExistsByUserIDAndType(userID uint, credentialType string) (bool, error)
}

// userCredentialRepository 用户凭证数据访问实现
type userCredentialRepository struct {
	db *gorm.DB
}

// NewUserCredentialRepository 创建用户凭证Repository
func NewUserCredentialRepository(db *gorm.DB) UserCredentialRepository {
	return &userCredentialRepository{db: db}
}

func (r *userCredentialRepository) Create(credential *table.UserCredential) error {
	return r.db.Create(credential).Error
}

func (r *userCredentialRepository) GetByID(id uint) (*table.UserCredential, error) {
	var credential table.UserCredential
	err := r.db.First(&credential, id).Error
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func (r *userCredentialRepository) GetByUserIDAndType(userID uint, credentialType string) (*table.UserCredential, error) {
	var credential table.UserCredential
	err := r.db.Where("user_id = ? AND credential_type = ?", userID, credentialType).First(&credential).Error
	if err != nil {
		return nil, err
	}
	return &credential, nil
}

func (r *userCredentialRepository) Update(credential *table.UserCredential) error {
	return r.db.Save(credential).Error
}

func (r *userCredentialRepository) Delete(id uint) error {
	return r.db.Delete(&table.UserCredential{}, id).Error
}

func (r *userCredentialRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&table.UserCredential{}).Error
}

func (r *userCredentialRepository) ExistsByUserIDAndType(userID uint, credentialType string) (bool, error) {
	var count int64
	err := r.db.Model(&table.UserCredential{}).Where("user_id = ? AND credential_type = ?", userID, credentialType).Count(&count).Error
	return count > 0, err
}
