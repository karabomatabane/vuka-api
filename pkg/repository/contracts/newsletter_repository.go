package contracts

import "vuka-api/pkg/models/db"

type NewsletterRepository interface {
	CreateSubscriber(subscriber *db.NewsletterSubscriber) error
	GetAllSubscribers() ([]db.NewsletterSubscriber, error)
	GetSubscriberByID(id string) (*db.NewsletterSubscriber, error)
	UpdateSubscriber(subscriber *db.NewsletterSubscriber) error
	DeleteSubscriber(id string) error
}
