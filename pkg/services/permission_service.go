package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
)

type PermissionService struct {
	repos *repository.Repositories
}

func NewPermissionService(repos *repository.Repositories) *PermissionService {
	return &PermissionService{repos: repos}
}

// Permission CRUD
func (s *PermissionService) CreatePermission(permission *db.Permission) error {
	return s.repos.Permission.CreatePermission(permission)
}

func (s *PermissionService) GetPermissionByID(id string) (*db.Permission, error) {
	permissionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.Permission.GetPermissionByID(permissionID)
}

func (s *PermissionService) GetAllPermissions() ([]db.Permission, error) {
	return s.repos.Permission.GetAllPermissions()
}

func (s *PermissionService) UpdatePermission(permission *db.Permission) error {
	return s.repos.Permission.UpdatePermission(permission)
}

func (s *PermissionService) DeletePermission(id string) error {
	permissionID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repos.Permission.DeletePermission(permissionID)
}

// Section CRUD
func (s *PermissionService) CreateSection(section *db.Section) error {
	return s.repos.Permission.CreateSection(section)
}

func (s *PermissionService) GetSectionByID(id string) (*db.Section, error) {
	sectionID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.Permission.GetSectionByID(sectionID)
}

func (s *PermissionService) GetAllSections() ([]db.Section, error) {
	return s.repos.Permission.GetAllSections()
}

func (s *PermissionService) UpdateSection(section *db.Section) error {
	return s.repos.Permission.UpdateSection(section)
}

func (s *PermissionService) DeleteSection(id string) error {
	sectionID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repos.Permission.DeleteSection(sectionID)
}

// RoleSectionPermission management
func (s *PermissionService) AssignPermissionToRole(roleSectionPermission *db.RoleSectionPermission) error {
	return s.repos.Permission.AssignPermissionToRole(roleSectionPermission)
}

func (s *PermissionService) RemovePermissionFromRole(roleID, sectionID, permissionID string) error {
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
	return s.repos.Permission.RemovePermissionFromRole(roleUUID, sectionUUID, permissionUUID)
}

func (s *PermissionService) GetRolePermissions(roleID string) ([]db.RoleSectionPermission, error) {
	roleUUID, err := uuid.Parse(roleID)
	if err != nil {
		return nil, err
	}
	return s.repos.Permission.GetRolePermissions(roleUUID)
}
