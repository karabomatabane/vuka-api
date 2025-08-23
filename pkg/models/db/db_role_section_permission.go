package db

import "github.com/google/uuid"

type RoleSectionPermission struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	SectionID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	PermissionID uuid.UUID `gorm:"type:uuid;primaryKey"`

	Role       Role
	Section    Section
	Permission Permission
}
