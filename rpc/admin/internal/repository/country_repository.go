package repository

import (
	"akatm/rpc/admin/orm/table"

	"gorm.io/gorm"
)

// CountryRepository 国家数据访问接口
type CountryRepository interface {
	// 创建国家
	Create(country *table.Country) error
	// 根据ID获取国家
	GetByID(id uint) (*table.Country, error)
	// 更新国家
	Update(country *table.Country) error
	// 删除国家
	Delete(id uint) error
	// 获取国家列表（分页）
	List(page, pageSize int64, keyword string, status int8) ([]*table.Country, int64, error)
	// 根据代码获取国家
	GetByCode(code string) (*table.Country, error)
	// 更新国家状态
	UpdateStatus(id uint, status int8) error
}

// countryRepository 国家数据访问实现
type countryRepository struct {
	db *gorm.DB
}

// NewCountryRepository 创建国家Repository
func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepository{db: db}
}

func (r *countryRepository) Create(country *table.Country) error {
	return r.db.Create(country).Error
}

func (r *countryRepository) GetByID(id uint) (*table.Country, error) {
	var country table.Country
	err := r.db.First(&country, id).Error
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) Update(country *table.Country) error {
	return r.db.Save(country).Error
}

func (r *countryRepository) Delete(id uint) error {
	return r.db.Delete(&table.Country{}, id).Error
}

func (r *countryRepository) List(page, pageSize int64, keyword string, status int8) ([]*table.Country, int64, error) {
	var countries []*table.Country
	var total int64

	// 构建查询
	query := r.db.Model(&table.Country{})

	// 状态筛选
	if status >= 0 {
		query = query.Where("status = ?", status)
	}

	// 关键词搜索
	if keyword != "" {
		query = query.Where("name LIKE ? OR name_en LIKE ? OR code LIKE ?",
			"%"+keyword+"%",
			"%"+keyword+"%",
			"%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Order("sort ASC, id ASC").Offset(int(offset)).Limit(int(pageSize)).Find(&countries).Error; err != nil {
		return nil, 0, err
	}

	return countries, total, nil
}

func (r *countryRepository) GetByCode(code string) (*table.Country, error) {
	var country table.Country
	err := r.db.Where("code = ?", code).First(&country).Error
	if err != nil {
		return nil, err
	}
	return &country, nil
}

func (r *countryRepository) UpdateStatus(id uint, status int8) error {
	return r.db.Model(&table.Country{}).Where("id = ?", id).Update("status", status).Error
}
