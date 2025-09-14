package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) contracts.ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(article *db.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) CreateWithTransaction(tx *gorm.DB, article *db.Article) error {
	return tx.Create(article).Error
}

func (r *articleRepository) CreateWithAssociations(article *db.Article) error {
	// Create the article with associations
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Create(article).Error
}

func (r *articleRepository) CreateWithAssociationsAndTransaction(tx *gorm.DB, article *db.Article) error {
	// Create the article with associations within a transaction
	return tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(article).Error
}

func (r *articleRepository) GetByID(id uuid.UUID) (*db.Article, error) {
	var article db.Article
	err := r.db.First(&article, id).Error
	return &article, err
}

func (r *articleRepository) GetAll() ([]db.Article, error) {
	var articles []db.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) Update(id uuid.UUID, updates map[string]any) error {
	return r.db.Model(&db.Article{}).Where("id = ?", id).Updates(updates).Error
}

func (r *articleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&db.Article{}, id).Error
}

func (r *articleRepository) GetByTitle(title string) (*db.Article, error) {
	var article db.Article
	err := r.db.Preload("Source").
		Preload("Region").
		Preload("Categories").
		Where("title = ?", title).
		First(&article).Error
	return &article, err
}

func (r *articleRepository) GetByOriginalUrl(url string) (*db.Article, error) {
	var article db.Article
	err := r.db.Preload("Source").
		Preload("Region").
		Preload("Categories").
		Where("original_url = ?", url).
		First(&article).Error
	return &article, err
}

func (r *articleRepository) ExistsByOriginalUrl(url string) (bool, error) {
	var count int64
	err := r.db.Model(&db.Article{}).Where("original_url = ?", url).Count(&count).Error
	return count > 0, err
}

func (r *articleRepository) GetWithRelations(id uuid.UUID) (*db.Article, error) {
	var article db.Article
	err := r.db.Preload("Source").
		Preload("Region").
		Preload("Categories").
		First(&article, id).Error
	return &article, err
}

func (r *articleRepository) GetAllWithRelations() ([]db.Article, error) {
	var articles []db.Article
	err := r.db.Preload("Genres").
		Preload("Directors").
		Preload("Cast").
		Preload("Reviews").
		Preload("Reviews.User").
		Find(&articles).Error
	return articles, err
}
