package contracts

import (
	"github.com/google/uuid"
	"vuka-api/pkg/models/db"
)

type RoleRepository interface {
	Create(role *db.Role) error
	Update(role *db.Role) error
	GetById(id string) (*db.Role, error)
	GetWithPermissions(id string) (*db.Role, error)
	GetAll() ([]db.Role, error)
	Delete(id uuid.UUID) error
}
