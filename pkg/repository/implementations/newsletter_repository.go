package implementations

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository/contracts"

	"gorm.io/gorm"
)

type NewsletterRepository struct {
	Db *gorm.DB
}

func NewNewsletterRepository(db *gorm.DB) contracts.NewsletterRepository {
	return &NewsletterRepository{Db: db}
}

func (r *NewsletterRepository) CreateSubscriber(subscriber *db.NewsletterSubscriber) error {
	return r.Db.Create(subscriber).Error
}

func (r *NewsletterRepository) GetAllSubscribers() ([]db.NewsletterSubscriber, error) {
	var subscribers []db.NewsletterSubscriber
	err := r.Db.Find(&subscribers).Error
	return subscribers, err
}

func (r *NewsletterRepository) GetSubscriberByID(id string) (*db.NewsletterSubscriber, error) {
	var subscriber db.NewsletterSubscriber
	err := r.Db.Where("id = ?", id).First(&subscriber).Error
	return &subscriber, err
}

func (r *NewsletterRepository) UpdateSubscriber(subscriber *db.NewsletterSubscriber) error {
	return r.Db.Save(subscriber).Error
}

func (r *NewsletterRepository) DeleteSubscriber(id string) error {
	return r.Db.Delete(&db.NewsletterSubscriber{}, "id = ?", id).Error
}
