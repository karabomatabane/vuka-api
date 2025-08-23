package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type castMemberRepository struct {
	db *gorm.DB
}

func NewCastMemberRepository(db *gorm.DB) contracts.CastMemberRepository {
	return &castMemberRepository{db: db}
}

func (r *castMemberRepository) Create(castMember *db.CastMember) error {
	return r.db.Create(castMember).Error
}

func (r *castMemberRepository) CreateBatch(castMembers []db.CastMember) error {
	if len(castMembers) == 0 {
		return nil
	}
	return r.db.Create(&castMembers).Error
}

func (r *castMemberRepository) GetByFilmID(filmID uuid.UUID) ([]db.CastMember, error) {
	var castMembers []db.CastMember
	err := r.db.Where("film_id = ?", filmID).Find(&castMembers).Error
	return castMembers, err
}

func (r *castMemberRepository) FindOrCreate(castMember *db.CastMember) (*db.CastMember, error) {
	err := r.db.Where("film_id = ? AND name = ?", castMember.FilmID, castMember.Name).FirstOrCreate(castMember).Error
	return castMember, err
}

func (r *castMemberRepository) CreateBatchWithTransaction(tx *gorm.DB, castMembers []db.CastMember) error {
	if len(castMembers) == 0 {
		return nil
	}

	// Create each cast member individually to handle duplicates
	for i := range castMembers {
		var existing db.CastMember
		err := tx.Where("film_id = ? AND name = ?", castMembers[i].FilmID, castMembers[i].Name).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			// Cast member doesn't exist, create it
			if err := tx.Create(&castMembers[i]).Error; err != nil {
				return err
			}
		} else if err != nil {
			// Some other error occurred
			return err
		}
		// If no error, the cast member already exists, so we skip it
	}

	return nil
}
