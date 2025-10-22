package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type directoryRepository struct {
	db *gorm.DB
}

func NewDirectoryRepository(db *gorm.DB) contracts.DirectoryRepository {
	return &directoryRepository{db: db}
}

func (r *directoryRepository) CreateCategory(category *db.DirectoryCategory) (error) {
	return r.db.Create(category).Error
}

func (r *directoryRepository) CreateEntry(entry *db.DirectoryEntry) (error) {
	return r.db.Create(entry).Error
}

func (r *directoryRepository) GetCategories() ([]db.DirectoryCategory, error) {
	var categories []db.DirectoryCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *directoryRepository) GetDirectoryEntriesByCategoryID(categoryID uuid.UUID) ([]db.DirectoryEntry, error) {
	var entries []db.DirectoryEntry
	err := r.db.Where("category_id = ?", categoryID).Find(&entries).Error
	return entries, err
}