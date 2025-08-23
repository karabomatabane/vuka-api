package user

import "github.com/google/uuid"

type UpdateUserRoleBody struct {
	UserID string    `json:"userId" validate:"required"`
	RoleID uuid.UUID `json:"roleId" validate:"required"`
}
