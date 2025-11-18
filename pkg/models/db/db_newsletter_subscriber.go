package db

type NewsletterSubscriber struct {
	Model
	PreferredName string `json:"preferredName"`
	Email         string `json:"email" gorm:"uniqueIndex"`
	PhoneNumber   string `json:"phoneNumber"`
}
