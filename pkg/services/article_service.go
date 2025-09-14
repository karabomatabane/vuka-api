package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
)

type ArticleService struct {
	repos *repository.Repositories
}

func NewArticleService(repos *repository.Repositories) *ArticleService {
	return &ArticleService{repos: repos}
}

func (s *ArticleService) CreateArticle(article *db.Article) error {
	return s.repos.Article.Create(article)
}

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

func (s *ArticleService) GetArticleByID(id string) (*db.Article, error) {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.Article.GetWithRelations(articleId)
}

func (s *ArticleService) GetAllArticles() ([]db.Article, error) {
	return s.repos.Article.GetAll()
}

func (s *ArticleService) UpdateArticle(id string, updates map[string]any) (*db.Article, error) {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	delete(updates, "id")
	err = s.repos.Article.Update(articleId, updates)
	if err != nil {
		return nil, err
	}
	return s.repos.Article.GetByID(articleId)
}

func (s *ArticleService) DeleteArticle(id string) error {
	articleId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repos.Article.Delete(articleId)
}
