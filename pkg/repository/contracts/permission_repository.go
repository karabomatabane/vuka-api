package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
)

type PermissionRepository interface {
	// Permission CRUD
	CreatePermission(permission *db.Permission) error
	GetPermissionByID(id uuid.UUID) (*db.Permission, error)
	GetAllPermissions() ([]db.Permission, error)
	UpdatePermission(permission *db.Permission) error
	DeletePermission(id uuid.UUID) error

	// Section CRUD
	CreateSection(section *db.Section) error
	GetSectionByID(id uuid.UUID) (*db.Section, error)
	GetAllSections() ([]db.Section, error)
	UpdateSection(section *db.Section) error
	DeleteSection(id uuid.UUID) error

	// RoleSectionPermission management
	AssignPermissionToRole(roleSectionPermission *db.RoleSectionPermission) error
	RemovePermissionFromRole(roleID, sectionID, permissionID uuid.UUID) error
	GetRolePermissions(roleID uuid.UUID) ([]db.RoleSectionPermission, error)
}
