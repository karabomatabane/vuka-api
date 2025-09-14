package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"gorm.io/gorm"
)

type sourceRepository struct {
	db *gorm.DB
}

func NewSourceRepository(db *gorm.DB) contracts.SourceRepository {
	return &sourceRepository{db: db}
}

func (r *sourceRepository) Create(source *db.Source) error {
	return r.db.Create(source).Error
}

func (r *sourceRepository) GetByID(id string) (*db.Source, error) {
	var source db.Source
	err := r.db.First(&source, "id = ?", id).Error
	return &source, err
}

func (r *sourceRepository) GetAll() ([]db.Source, error) {
	var sources []db.Source
	err := r.db.Find(&sources).Error
	return sources, err
}

func (r *sourceRepository) Update(source *db.Source) error {
	return r.db.Save(source).Error
}

func (r *sourceRepository) Delete(id string) error {
	return r.db.Delete(&db.Source{}, "id = ?", id).Error
}
