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

func (s *ArticleService) GetArticleByID(id uuid.UUID) (*db.Article, error) {
	return s.repos.Article.GetWithRelations(id)
}

func (s *ArticleService) GetAllArticles() ([]db.Article, error) {
	return s.repos.Article.GetAll()
}

func (s *ArticleService) UpdateArticle(article *db.Article) error {
	return s.repos.Article.Update(article)
}

func (s *ArticleService) DeleteArticle(id uuid.UUID) error {
	return s.repos.Article.Delete(id)
}
