package repository

import (
	"akatm/rpc/iam/orm/table"

	"gorm.io/gorm"
)

// UserSessionRepository 用户会话数据访问接口
type UserSessionRepository interface {
	// 创建会话
	Create(session *table.UserSession) error
	// 根据Token获取会话
	GetByToken(token string) (*table.UserSession, error)
	// 根据用户ID获取会话列表
	GetByUserID(userID uint) ([]*table.UserSession, error)
	// 更新会话
	Update(session *table.UserSession) error
	// 删除会话
	Delete(token string) error
	// 删除用户的所有会话
	DeleteByUserID(userID uint) error
	// 检查Token是否存在
	ExistsByToken(token string) (bool, error)
	// 清理过期会话
	CleanExpiredSessions() error
}

// userSessionRepository 用户会话数据访问实现
type userSessionRepository struct {
	db *gorm.DB
}

// NewUserSessionRepository 创建用户会话Repository
func NewUserSessionRepository(db *gorm.DB) UserSessionRepository {
	return &userSessionRepository{db: db}
}

func (r *userSessionRepository) Create(session *table.UserSession) error {
	return r.db.Create(session).Error
}

func (r *userSessionRepository) GetByToken(token string) (*table.UserSession, error) {
	var session table.UserSession
	err := r.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *userSessionRepository) GetByUserID(userID uint) ([]*table.UserSession, error) {
	var sessions []*table.UserSession
	err := r.db.Where("user_id = ?", userID).Find(&sessions).Error
	return sessions, err
}

func (r *userSessionRepository) Update(session *table.UserSession) error {
	return r.db.Save(session).Error
}

func (r *userSessionRepository) Delete(token string) error {
	return r.db.Where("token = ?", token).Delete(&table.UserSession{}).Error
}

func (r *userSessionRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&table.UserSession{}).Error
}

func (r *userSessionRepository) ExistsByToken(token string) (bool, error) {
	var count int64
	err := r.db.Model(&table.UserSession{}).Where("token = ?", token).Count(&count).Error
	return count > 0, err
}

func (r *userSessionRepository) CleanExpiredSessions() error {
	// 删除过期的会话
	return r.db.Where("expires_at < ?", "UNIX_TIMESTAMP()").Delete(&table.UserSession{}).Error
}
