package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type filmRepository struct {
	db *gorm.DB
}

func NewFilmRepository(db *gorm.DB) contracts.FilmRepository {
	return &filmRepository{db: db}
}

func (r *filmRepository) Create(film *db.Film) error {
	// Use FullSaveAssociations to save many-to-many relationships
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(film).Error
}

func (r *filmRepository) CreateWithAssociations(film *db.Film) error {
	// Create the film with associations
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(film).Error
}

func (r *filmRepository) CreateWithAssociationsAndTransaction(tx *gorm.DB, film *db.Film) error {
	// Create the film with associations within a transaction
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(film).Error
}

func (r *filmRepository) GetByID(id uuid.UUID) (*db.Film, error) {
	var film db.Film
	err := r.db.First(&film, id).Error
	return &film, err
}

func (r *filmRepository) GetAll() ([]db.Film, error) {
	var films []db.Film
	err := r.db.Find(&films).Error
	return films, err
}

func (r *filmRepository) Update(film *db.Film) error {
	return r.db.Save(film).Error
}

func (r *filmRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&db.Film{}, id).Error
}

func (r *filmRepository) GetByName(name string) (*db.Film, error) {
	var film db.Film
	err := r.db.Where("name = ?", name).First(&film).Error
	return &film, err
}

func (r *filmRepository) GetWithRelations(id uuid.UUID) (*db.Film, error) {
	var film db.Film
	err := r.db.Preload("Genres").
		Preload("Directors").
		Preload("Cast").
		Preload("Reviews").
		Preload("Reviews.User").
		First(&film, id).Error
	return &film, err
}

func (r *filmRepository) GetAllWithRelations() ([]db.Film, error) {
	var films []db.Film
	err := r.db.Preload("Genres").
		Preload("Directors").
		Preload("Cast").
		Preload("Reviews").
		Preload("Reviews.User").
		Find(&films).Error
	return films, err
}

func (r *filmRepository) CreateWithTransaction(tx *gorm.DB, film *db.Film) error {
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(film).Error
}
