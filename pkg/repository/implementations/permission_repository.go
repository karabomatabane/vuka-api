package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) contracts.PermissionRepository {
	return &permissionRepository{db: db}
}

// Permission CRUD
func (r *permissionRepository) CreatePermission(permission *db.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) GetPermissionByID(id uuid.UUID) (*db.Permission, error) {
	var permission db.Permission
	err := r.db.First(&permission, id).Error
	return &permission, err
}

func (r *permissionRepository) GetAllPermissions() ([]db.Permission, error) {
	var permissions []db.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) UpdatePermission(permission *db.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) DeletePermission(id uuid.UUID) error {
	return r.db.Delete(&db.Permission{}, id).Error
}

// Section CRUD
func (r *permissionRepository) CreateSection(section *db.Section) error {
	return r.db.Create(section).Error
}

func (r *permissionRepository) GetSectionByID(id uuid.UUID) (*db.Section, error) {
	var section db.Section
	err := r.db.First(&section, id).Error
	return &section, err
}

func (r *permissionRepository) GetAllSections() ([]db.Section, error) {
	var sections []db.Section
	err := r.db.Find(&sections).Error
	return sections, err
}

func (r *permissionRepository) UpdateSection(section *db.Section) error {
	return r.db.Save(section).Error
}

func (r *permissionRepository) DeleteSection(id uuid.UUID) error {
	return r.db.Delete(&db.Section{}, id).Error
}

// RoleSectionPermission management
func (r *permissionRepository) AssignPermissionToRole(roleSectionPermission *db.RoleSectionPermission) error {
	return r.db.Create(roleSectionPermission).Error
}

func (r *permissionRepository) RemovePermissionFromRole(roleID, sectionID, permissionID uuid.UUID) error {
	return r.db.Where("role_id = ? AND section_id = ? AND permission_id = ?",
		roleID, sectionID, permissionID).
		Delete(&db.RoleSectionPermission{}).Error
}

func (r *permissionRepository) GetRolePermissions(roleID uuid.UUID) ([]db.RoleSectionPermission, error) {
	var permissions []db.RoleSectionPermission
	err := r.db.Where("role_id = ?", roleID).
		Preload("Permission").
		Preload("Section").
		Find(&permissions).Error
	return permissions, err
}
