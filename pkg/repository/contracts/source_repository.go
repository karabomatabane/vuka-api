package contracts

import "vuka-api/pkg/models/db"

type SourceRepository interface {
	Create(source *db.Source) error
	GetByID(id string) (*db.Source, error)
	GetAll() ([]db.Source, error)
	Update(source *db.Source) error
	Delete(id string) error
}
