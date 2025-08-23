package db

import "github.com/google/uuid"

type User struct {
	Model
	Username     string    `json:"username"`
	PasswordHash string    `json:"password"`
	RoleID       uuid.UUID `json:"roleId"`
	Role         Role      `json:"role" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
