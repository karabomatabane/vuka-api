package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"gorm.io/gorm"
)

type directorRepository struct {
	db *gorm.DB
}

func NewDirectorRepository(db *gorm.DB) contracts.DirectorRepository {
	return &directorRepository{db: db}
}

func (r *directorRepository) FindOrCreate(name string) (*db.Director, error) {
	var director db.Director
	err := r.db.Where("name = ?", name).FirstOrCreate(&director, db.Director{Name: name}).Error
	return &director, err
}

func (r *directorRepository) Update(name, imgUrl string) (*db.Director, error) {
	director := db.Director{}
	err := r.db.Model(&director).Where("name = ?", name).Update("img_url", imgUrl).Error
	err = r.db.Where("name = ?", name).First(&director).Error
	return &director, err
}

func (r *directorRepository) GetAll() ([]db.Director, error) {
	var directors []db.Director
	err := r.db.Find(&directors).Error
	return directors, err
}

func (r *directorRepository) GetByName(name string) (*db.Director, error) {
	var director db.Director
	err := r.db.Where("name = ?", name).First(&director).Error
	return &director, err
}

func (r *directorRepository) FindOrCreateWithTransaction(tx *gorm.DB, name string) (*db.Director, error) {
	var director db.Director
	err := tx.Where("name = ?", name).FirstOrCreate(&director, db.Director{Name: name}).Error
	return &director, err
}
