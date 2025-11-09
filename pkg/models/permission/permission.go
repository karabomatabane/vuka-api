package permission

import "github.com/google/uuid"

type AssignPermissionRequest struct {
	RoleID       uuid.UUID `json:"roleId" validate:"required"`
	SectionID    uuid.UUID `json:"sectionId" validate:"required"`
	PermissionID uuid.UUID `json:"permissionId" validate:"required"`
}

type RemovePermissionRequest struct {
	RoleID       string `json:"roleId" validate:"required"`
	SectionID    string `json:"sectionId" validate:"required"`
	PermissionID string `json:"permissionId" validate:"required"`
}

type PermissionResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type SectionResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type RolePermissionResponse struct {
	Section    SectionResponse    `json:"section"`
	Permission PermissionResponse `json:"permission"`
}
