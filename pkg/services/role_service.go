package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
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

func (s *RoleService) Update(role *db.Role) error {
	return s.Repos.Role.Update(role)
}

func (s *RoleService) AssignPermissionToRole(roleSectionPermission *db.RoleSectionPermission) error {
	return s.Repos.Permission.AssignPermissionToRole(roleSectionPermission)
}

func (s *RoleService) RemovePermissionFromRole(roleID, sectionID, permissionID string) error {
	roleUUID, err := uuid.Parse(roleID)
	if err != nil {
		return err
	}
	sectionUUID, err := uuid.Parse(sectionID)
	if err != nil {
		return err
	}
	permissionUUID, err := uuid.Parse(permissionID)
	if err != nil {
		return err
	}
	return s.Repos.Permission.RemovePermissionFromRole(roleUUID, sectionUUID, permissionUUID)
}

func (s *RoleService) GetRolePermissions(roleID string) ([]db.RoleSectionPermission, error) {
	roleUUID, err := uuid.Parse(roleID)
	if err != nil {
		return nil, err
	}
	return s.Repos.Permission.GetRolePermissions(roleUUID)
}

func (s *RoleService) GetAll() ([]db.Role, error) {
	return s.Repos.Role.GetAll()
}

func (s *RoleService) GetById(id string) (*db.Role, error) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.Repos.Role.GetById(roleId)
}

func (s *RoleService) GetWithPermissions(id string) (*db.Role, error) {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.Repos.Role.GetWithPermissions(roleId)
}

func (s *RoleService) Delete(id string) error {
	roleId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.Repos.Role.Delete(roleId)
}
