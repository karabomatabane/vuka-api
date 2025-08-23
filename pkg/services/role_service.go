package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
)

type RoleService struct {
	Repos *repository.Repositories
}

func NewRoleService(repos *repository.Repositories) *RoleService {
	return &RoleService{Repos: repos}
}

func (s *RoleService) Create(role *db.Role) error {
	return s.Repos.Role.Create(role)
}
