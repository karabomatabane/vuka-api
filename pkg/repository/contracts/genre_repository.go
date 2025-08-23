package contracts

import (
	"vuka-api/pkg/models/db"

	"gorm.io/gorm"
)

type GenreRepository interface {
	FindOrCreate(name string) (*db.Genre, error)
	GetAll() ([]db.Genre, error)
	GetByName(name string) (*db.Genre, error)
	FindOrCreateWithTransaction(tx *gorm.DB, name string) (*db.Genre, error)
}
