package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
)

// DirectoryService is a service for managing directories.
type DirectoryService struct {
	repo contracts.DirectoryRepository
}

// NewDirectoryService creates a new DirectoryService.
func NewDirectoryService(repo contracts.DirectoryRepository) *DirectoryService {
	return &DirectoryService{repo: repo}
}

// GetAllDirectories returns all directories.
func (s *DirectoryService) GetAllDirectories() ([]db.DirectoryCategory, error) {
	return s.repo.GetCategories()
}

// CreateDirectoryCategory creates a new directory category.
func (s *DirectoryService) CreateDirectoryCategory(category *db.DirectoryCategory) error {
	return s.repo.CreateCategory(category)
}

// CreateDirectoryEntry creates a new directory entry.
func (s *DirectoryService) CreateDirectoryEntry(entry *db.DirectoryEntry) error {
	return s.repo.CreateEntry(entry)
}

// GetDirectoryEntriesByCategoryID returns all directory entries for a specific category.
func (s *DirectoryService) GetDirectoryEntriesByCategoryID(categoryID string) ([]db.DirectoryEntry, error) {
	uuidCategoryID, err := uuid.Parse(categoryID)
	if err != nil {
		return nil, err
	}
	entries, err := s.repo.GetDirectoryEntriesByCategoryID(uuidCategoryID)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
