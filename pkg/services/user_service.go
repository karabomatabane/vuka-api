package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repos *repository.Repositories
}

func NewUserService(repos *repository.Repositories) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) GetUserByID(id string) (*db.User, error) {
	filmId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.User.GetByID(filmId)
}

func (s *UserService) GetUserByUsername(username string) (*db.User, error) {
	return s.repos.User.GetByUsername(username)
}

func (s *UserService) GetAllUsers() ([]db.User, error) {
	return s.repos.User.GetAll()
}

// GetUUIDFromMongoID converts a MongoDB ObjectId to a deterministic UUID
// This is useful for maintaining consistency when referencing users from other data
func (s *UserService) GetUUIDFromMongoID(mongoID string) uuid.UUID {
	if mongoID == "" {
		return uuid.New()
	}

	// Try to parse as UUID first
	if userUUID, err := uuid.Parse(mongoID); err == nil {
		return userUUID
	}

	// Generate deterministic UUID from MongoDB ObjectId
	return uuid.NewSHA1(uuid.NameSpaceOID, []byte(mongoID))
}
