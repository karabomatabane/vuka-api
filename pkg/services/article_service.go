package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
)

// ArticleService ...
type ArticleService struct {
	repos *repository.Repositories
}

// NewArticleService ...
func NewArticleService(repos *repository.Repositories) *ArticleService {
	return &ArticleService{repos: repos}
}

// CreateArticle ...
func (s *ArticleService) CreateArticle(article *db.Article) error {
	return s.repos.Article.Create(article)
}

// CreateArticleIfNotExists ...
func (s *ArticleService) CreateArticleIfNotExists(article *db.Article) (bool, error) {
	// Check if article already exists by original URL
	exists, err := s.repos.Article.ExistsByOriginalUrl(article.OriginalUrl)
	if err != nil {
		return false, err
	}

	if exists {
		return false, nil // Article exists, not created
	}

	// Article doesn't exist, create it
	err = s.repos.Article.Create(article)
	if err != nil {
		return false, err
	}

	return true, nil // Article created successfully
}

// GetArticleByID ...
func (s *ArticleService) GetArticleByID(id string) (*db.Article, error) {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.Article.GetWithRelations(articleId)
}

// GetAllArticles ...
func (s *ArticleService) GetAllArticles() ([]db.Article, error) {
	return s.repos.Article.GetAllWithRelations()
}

// GetAllArticlesPaginated returns paginated articles with relations
func (s *ArticleService) GetAllArticlesPaginated(limit, offset int) ([]db.Article, int64, error) {
	return s.repos.Article.GetAllWithRelationsPaginated(limit, offset)
}

// UpdateArticle ...
func (s *ArticleService) UpdateArticle(id string, updates map[string]any) (*db.Article, error) {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if categoryIds, ok := updates["categoryIds"].([]any); ok {
		var categories []db.Category
		if err := s.repos.Category.FindIn("id", categoryIds, &categories); err != nil {
			return nil, err
		}

		article, err := s.repos.Article.GetByID(articleId)
		if err != nil {
			return nil, err
		}

		if err := s.repos.Article.SetCategories(article, categories); err != nil {
			return nil, err
		}

		delete(updates, "categoryIds")
	}

	delete(updates, "id")
	err = s.repos.Article.Update(articleId, updates)
	if err != nil {
		return nil, err
	}
	return s.repos.Article.GetByID(articleId)
}

// DeleteArticle ...
func (s *ArticleService) DeleteArticle(id string) error {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repos.Article.Delete(articleId)
}

// SetArticleCategories sets the categories for a given article.
func (s *ArticleService) SetArticleCategories(article *db.Article, categories []db.Category) error {
	return s.repos.Article.SetCategories(article, categories)
}
