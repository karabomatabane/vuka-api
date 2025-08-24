package contracts

import (
	"vuka-api/pkg/models/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *db.Article) error
	GetByID(id uuid.UUID) (*db.Article, error)
	GetAll() ([]db.Article, error)
	Update(id uuid.UUID, updates map[string]any) error
	Delete(id uuid.UUID) error
	GetByTitle(title string) (*db.Article, error)
	GetWithRelations(id uuid.UUID) (*db.Article, error)
	GetAllWithRelations() ([]db.Article, error)
	CreateWithTransaction(tx *gorm.DB, article *db.Article) error
	CreateWithAssociations(article *db.Article) error
	CreateWithAssociationsAndTransaction(tx *gorm.DB, article *db.Article) error
}
