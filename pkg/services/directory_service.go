package services

import (
	"vuka-api/pkg/models/db"
	models "vuka-api/pkg/models/directory"
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

// GetDirectoryOverview returns an overview of the directory.
func (s *DirectoryService) GetDirectoryOverview(userID uuid.UUID) (*models.DirectoryOverviewResponse, error) {
	categories, err := s.repo.GetCategories()
	if err != nil {
		return nil, err
	}

	respCategories := make([]models.DirectoryCategoryResponse, len(categories))
	for i, c := range categories {
		count, err := s.repo.CountEntriesByCategoryID(c.ID)
		if err != nil {
			return nil, err
		}
		respCategories[i] = models.DirectoryCategoryResponse{
			ID:           c.ID,
			Name:         c.Name,
			TotalEntries: count,
		}
	}

	pinned, err := s.repo.GetPinnedDirectories(userID)
	if err != nil {
		return nil, err
	}
	pinnedResp := make([]models.DirectoryCategoryResponse, len(pinned))
	for i, p := range pinned {
		pinnedResp[i] = models.DirectoryCategoryResponse{
			ID:   p.ID,
			Name: p.Name,
		}
	}

	recent, err := s.repo.GetRecentDirectories(userID)
	if err != nil {
		return nil, err
	}
	recentResp := make([]models.DirectoryCategoryResponse, len(recent))
	for i, r := range recent {
		recentResp[i] = models.DirectoryCategoryResponse{
			ID:   r.ID,
			Name: r.Name,
		}
	}

	return &models.DirectoryOverviewResponse{
		Categories: respCategories,
		Personalized: models.PersonalisedData{
			Pinned: pinnedResp,
			Recent: recentResp,
		},
	}, nil
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

// GetDirectoryEntriesByCategoryID returns a directory category with all its entries.
func (s *DirectoryService) GetDirectoryEntriesByCategoryID(categoryID string) (*db.DirectoryCategory, error) {
	uuidCategoryID, err := uuid.Parse(categoryID)
	if err != nil {
		return nil, err
	}
	category, err := s.repo.GetDirectoryEntriesByCategoryID(uuidCategoryID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *DirectoryService) GetDirectoryEntryByID(entryID string) (*db.DirectoryEntry, error) {
	uuidEntryID, err := uuid.Parse(entryID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetDirectoryEntryByID(uuidEntryID)
}
