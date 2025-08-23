package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user *db.User) error
	CreateBatch(users []db.User) error
	GetByID(id uuid.UUID) (*db.User, error)
	GetByEmail(email string) (*db.User, error)
	GetByUsername(username string) (*db.User, error)
	Update(user *db.User) error
	Delete(id uuid.UUID) error
	GetAll() ([]db.User, error)
}
