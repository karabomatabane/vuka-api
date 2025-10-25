package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
)

type DirectoryRepository interface {
	CreateCategory(category *db.DirectoryCategory) error
	GetCategories() ([]db.DirectoryCategory, error)
	CreateEntry(entry *db.DirectoryEntry) error
	GetDirectoryEntriesByCategoryID(categoryID uuid.UUID) ([]db.DirectoryEntry, error)
	CountEntriesByCategoryID(categoryID uuid.UUID) (int64, error)
	GetPinnedDirectories(userID uuid.UUID) ([]db.DirectoryCategory, error)
	GetRecentDirectories(userID uuid.UUID) ([]db.DirectoryCategory, error)
}
