package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmRepository interface {
	Create(film *db.Film) error
	GetByID(id uuid.UUID) (*db.Film, error)
	GetAll() ([]db.Film, error)
	Update(film *db.Film) error
	Delete(id uuid.UUID) error
	GetByName(name string) (*db.Film, error)
	GetWithRelations(id uuid.UUID) (*db.Film, error)
	GetAllWithRelations() ([]db.Film, error)
	CreateWithTransaction(tx *gorm.DB, film *db.Film) error
	CreateWithAssociations(film *db.Film) error
	CreateWithAssociationsAndTransaction(tx *gorm.DB, film *db.Film) error
}
