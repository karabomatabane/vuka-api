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

// GetAllArticlesPaginatedAndSearch returns paginated articles with relations and search
func (s *ArticleService) GetAllArticlesPaginatedAndSearch(limit, offset int, search string) ([]db.Article, int64, error) {
	return s.repos.Article.GetAllWithRelationsPaginatedAndSearch(limit, offset, search)
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

	// Remove fields that shouldn't be updated directly
	delete(updates, "id")
	delete(updates, "categories")
	delete(updates, "source")
	delete(updates, "region")
	delete(updates, "images")

	// Convert camelCase JSON field names to snake_case database column names
	if val, ok := updates["isFeatured"]; ok {
		updates["is_featured"] = val
		delete(updates, "isFeatured")
	}
	if val, ok := updates["contentBody"]; ok {
		updates["content_body"] = val
		delete(updates, "contentBody")
	}
	if val, ok := updates["originalUrl"]; ok {
		updates["original_url"] = val
		delete(updates, "originalUrl")
	}
	if val, ok := updates["publishedAt"]; ok {
		updates["published_at"] = val
		delete(updates, "publishedAt")
	}
	if val, ok := updates["sourceId"]; ok {
		updates["source_id"] = val
		delete(updates, "sourceId")
	}

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
