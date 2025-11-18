package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
)

type NewsletterService struct {
	repo *repository.Repositories
}

func NewNewsletterService(repo *repository.Repositories) *NewsletterService {
	return &NewsletterService{repo: repo}
}

func (s *NewsletterService) Subscribe(subscriber *db.NewsletterSubscriber) error {
	return s.repo.Newsletter.CreateSubscriber(subscriber)
}

func (s *NewsletterService) GetAllSubscribers() ([]db.NewsletterSubscriber, error) {
	return s.repo.Newsletter.GetAllSubscribers()
}

func (s *NewsletterService) GetSubscriberByID(id string) (*db.NewsletterSubscriber, error) {
	return s.repo.Newsletter.GetSubscriberByID(id)
}

func (s *NewsletterService) UpdateSubscriber(subscriber *db.NewsletterSubscriber) error {
	return s.repo.Newsletter.UpdateSubscriber(subscriber)
}

func (s *NewsletterService) DeleteSubscriber(id string) error {
	return s.repo.Newsletter.DeleteSubscriber(id)
}
