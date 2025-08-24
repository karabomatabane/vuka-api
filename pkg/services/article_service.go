package services

import (
	"github.com/google/uuid"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
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
