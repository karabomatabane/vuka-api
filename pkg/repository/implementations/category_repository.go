package implementations

import (
	"gorm.io/gorm"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) contracts.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() ([]db.Category, error) {
	var categories []db.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindIn(field string, values []interface{}, target interface{}) error {
	return r.db.Where(field+" IN ?", values).Find(target).Error
}

func (r *categoryRepository) FindByName(name string) (*db.Category, error) {
	var category db.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *db.Category) error {
	return r.db.Create(category).Error
}