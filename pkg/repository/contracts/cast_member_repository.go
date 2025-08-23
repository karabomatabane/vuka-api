package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CastMemberRepository interface {
	Create(castMember *db.CastMember) error
	CreateBatch(castMembers []db.CastMember) error
	GetByFilmID(filmID uuid.UUID) ([]db.CastMember, error)
	FindOrCreate(castMember *db.CastMember) (*db.CastMember, error)
	CreateBatchWithTransaction(tx *gorm.DB, castMembers []db.CastMember) error
}
