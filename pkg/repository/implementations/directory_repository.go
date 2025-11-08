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

func (r *directoryRepository) CreateCategory(category *db.DirectoryCategory) error {
	return r.db.Create(category).Error
}

func (r *directoryRepository) CreateEntry(entry *db.DirectoryEntry) error {
	// GORM automatically handles the nested creation of ContactInfo
	// due to the foreign key relationship defined in the model
	return r.db.Create(entry).Error
}

func (r *directoryRepository) GetCategories() ([]db.DirectoryCategory, error) {
	var categories []db.DirectoryCategory
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *directoryRepository) GetDirectoryEntriesByCategoryID(categoryID uuid.UUID) (*db.DirectoryCategory, error) {
	var category db.DirectoryCategory
	err := r.db.Preload("Directories.ContactInfo").First(&category, categoryID).Error
	return &category, err
}

// GetDirectoryEntryByID implements contracts.DirectoryRepository.
func (r *directoryRepository) GetDirectoryEntryByID(entryID uuid.UUID) (*db.DirectoryEntry, error) {
	var entry db.DirectoryEntry
	err := r.db.Preload("ContactInfo").Preload("Category").First(&entry, entryID).Error
	return &entry, err
}

func (r *directoryRepository) CountEntriesByCategoryID(categoryID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&db.DirectoryEntry{}).Where("category_id = ?", categoryID).Count(&count).Error
	return count, err
}

func (r *directoryRepository) GetPinnedDirectories(userID uuid.UUID) ([]db.DirectoryCategory, error) {
	var categories []db.DirectoryCategory
	err := r.db.
		Joins("JOIN user_directory_meta ON user_directory_meta.directory_id = directory_categories.id").
		Where("user_directory_meta.user_id = ? AND user_directory_meta.pinned = ?", userID, true).
		Find(&categories).Error
	return categories, err
}

func (r *directoryRepository) GetRecentDirectories(userID uuid.UUID) ([]db.DirectoryCategory, error) {
	var categories []db.DirectoryCategory
	err := r.db.
		Joins("JOIN user_directory_meta ON user_directory_meta.directory_id = directory_categories.id").
		Where("user_directory_meta.user_id = ?", userID).
		Order("user_directory_meta.last_accessed desc").
		Limit(5).
		Find(&categories).Error
	return categories, err
}
