package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
)

type SourceService struct {
	repos *repository.Repositories
}

func NewSourceService(repos *repository.Repositories) *SourceService {
	return &SourceService{repos: repos}
}

func (s *SourceService) CreateSource(source *db.Source) error {
	return s.repos.Source.Create(source)
}

func (s *SourceService) GetSourceByID(id string) (*db.Source, error) {
	return s.repos.Source.GetByID(id)
}

func (s *SourceService) GetAllSources() ([]db.Source, error) {
	return s.repos.Source.GetAll()
}

func (s *SourceService) UpdateSource(source *db.Source) error {
	return s.repos.Source.Update(source)
}

func (s *SourceService) DeleteSource(id string) error {
	return s.repos.Source.Delete(id)
}
