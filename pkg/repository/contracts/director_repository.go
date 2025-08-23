package contracts

import (
	"vuka-api/pkg/models/db"

	"gorm.io/gorm"
)

type DirectorRepository interface {
	FindOrCreate(name string) (*db.Director, error)
	Update(name, imgUrl string) (*db.Director, error)
	GetAll() ([]db.Director, error)
	GetByName(name string) (*db.Director, error)
	FindOrCreateWithTransaction(tx *gorm.DB, name string) (*db.Director, error)
}
