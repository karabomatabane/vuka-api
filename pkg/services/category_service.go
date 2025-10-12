package services

import (
	"errors"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"gorm.io/gorm"
)

// CategoryService is a service for managing categories.
type CategoryService struct {
	repo contracts.CategoryRepository
}

// NewCategoryService creates a new CategoryService.
func NewCategoryService(repo contracts.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAllCategories returns all categories.
func (s *CategoryService) GetAllCategories() ([]db.Category, error) {
	return s.repo.FindAll()
}

// CreateCategory creates a new category.
func (s *CategoryService) CreateCategory(category *db.Category) error {
	return s.repo.Create(category)
}

// FindOrCreate finds a category by name or creates it if it doesn't exist.
func (s *CategoryService) FindOrCreate(name string) (*db.Category, error) {
	category, err := s.repo.FindByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Category not found, create it
			newCategory := &db.Category{Name: name}
			err = s.repo.Create(newCategory)
			if err != nil {
				return nil, err
			}
			return newCategory, nil
		}
		return nil, err
	}
	return category, nil
}
