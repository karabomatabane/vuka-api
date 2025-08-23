package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) contracts.ReviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(review *db.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) CreateBatch(reviews []db.Review) error {
	if len(reviews) == 0 {
		return nil
	}
	return r.db.Create(&reviews).Error
}

func (r *reviewRepository) GetByFilmID(filmID uuid.UUID) ([]db.Review, error) {
	var reviews []db.Review
	err := r.db.Where("film_id = ?", filmID).Preload("User").Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetByUserID(userID uuid.UUID) ([]db.Review, error) {
	var reviews []db.Review
	err := r.db.Where("user_id = ?", userID).Preload("Film").Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) Update(review *db.Review) error {
	return r.db.Save(review).Error
}

func (r *reviewRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&db.Review{}, id).Error
}

func (r *reviewRepository) CreateBatchWithTransaction(tx *gorm.DB, reviews []db.Review) error {
	if len(reviews) == 0 {
		return nil
	}
	return tx.Create(&reviews).Error
}
