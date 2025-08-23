package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) contracts.GenreRepository {
	return &genreRepository{db: db}
}

func (r *genreRepository) FindOrCreate(name string) (*db.Genre, error) {
	var genre db.Genre
	err := r.db.Where("name = ?", name).FirstOrCreate(&genre, db.Genre{Name: name}).Error
	return &genre, err
}

func (r *genreRepository) GetAll() ([]db.Genre, error) {
	var genres []db.Genre
	err := r.db.Find(&genres).Error
	return genres, err
}

func (r *genreRepository) GetByName(name string) (*db.Genre, error) {
	var genre db.Genre
	err := r.db.Where("name = ?", name).First(&genre).Error
	return &genre, err
}

func (r *genreRepository) FindOrCreateWithTransaction(tx *gorm.DB, name string) (*db.Genre, error) {
	var genre db.Genre
	err := tx.Where("name = ?", name).FirstOrCreate(&genre, db.Genre{Name: name}).Error
	return &genre, err
}
