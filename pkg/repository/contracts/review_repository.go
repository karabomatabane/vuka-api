package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *db.Review) error
	CreateBatch(reviews []db.Review) error
	GetByFilmID(filmID uuid.UUID) ([]db.Review, error)
	GetByUserID(userID uuid.UUID) ([]db.Review, error)
	Update(review *db.Review) error
	Delete(id uuid.UUID) error
	CreateBatchWithTransaction(tx *gorm.DB, reviews []db.Review) error
}
