package services

import (
	"fmt"
	"log"
	"time"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
)

type NewsletterService struct {
	repo         *repository.Repositories
	emailService *EmailService
	articleRepo  *repository.Repositories
}

func NewNewsletterService(repo *repository.Repositories) *NewsletterService {
	return &NewsletterService{
		repo:         repo,
		emailService: NewEmailService(),
		articleRepo:  repo,
	}
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

// SendNewsletter sends newsletter to all subscribers
func (s *NewsletterService) SendNewsletter(subject, content string, useTemplate bool, templateData map[string]interface{}) error {
	subscribers, err := s.GetAllSubscribers()
	if err != nil {
		return fmt.Errorf("failed to get subscribers: %w", err)
	}

	if len(subscribers) == 0 {
		return fmt.Errorf("no subscribers found")
	}

	log.Printf("Preparing to send newsletter to %d subscribers", len(subscribers))

	// Prepare email data for all subscribers
	var emailDataList []EmailData
	for _, subscriber := range subscribers {
		emailData := EmailData{
			ToEmail: subscriber.Email,
			ToName:  subscriber.PreferredName,
			Subject: subject,
		}

		if useTemplate {
			emailData.TemplateName = "newsletter"
			if templateData == nil {
				templateData = make(map[string]interface{})
			}
			templateData["SubscriberName"] = subscriber.PreferredName
			emailData.TemplateData = templateData
		} else {
			emailData.PlainTextBody = content
		}

		emailDataList = append(emailDataList, emailData)
	}

	// Send bulk emails
	errors := s.emailService.SendBulkEmail(emailDataList)
	if len(errors) > 0 {
		return fmt.Errorf("failed to send some emails: %d errors occurred", len(errors))
	}

	log.Printf("Newsletter sent successfully to all %d subscribers", len(subscribers))
	return nil
}

// SendNewsletterWithLatestArticles sends newsletter with the latest featured articles
func (s *NewsletterService) SendNewsletterWithLatestArticles(subject string, limit int) error {
	// Get latest featured articles
	articles, err := s.articleRepo.Article.GetAllWithRelations()
	if err != nil {
		return fmt.Errorf("failed to get articles: %w", err)
	}

	// Filter for featured articles and limit
	var featuredArticles []db.Article
	for _, article := range articles {
		if article.IsFeatured && len(featuredArticles) < limit {
			featuredArticles = append(featuredArticles, article)
		}
	}

	if len(featuredArticles) == 0 {
		return fmt.Errorf("no featured articles found")
	}

	// Prepare template data
	templateData := map[string]interface{}{
		"Articles": featuredArticles,
		"Date":     time.Now().Format("January 2, 2006"),
		"Year":     time.Now().Year(),
	}

	return s.SendNewsletter(subject, "", true, templateData)
}

// SendTestEmail sends a test email to verify SMTP configuration
func (s *NewsletterService) SendTestEmail(toEmail, toName string) error {
	emailData := EmailData{
		ToEmail:       toEmail,
		ToName:        toName,
		Subject:       "Test Email from Vuka Newsletter",
		PlainTextBody: fmt.Sprintf("Hello %s,\n\nThis is a test email to verify your SMTP configuration is working correctly.\n\nBest regards,\nVuka Team", toName),
	}

	return s.emailService.SendEmail(emailData)
}

// TestSMTPConnection tests the SMTP connection
func (s *NewsletterService) TestSMTPConnection() error {
	return s.emailService.TestConnection()
}

// GenerateNewsletterPreview generates HTML preview of the newsletter
func (s *NewsletterService) GenerateNewsletterPreview(limit int, customData map[string]interface{}) (string, error) {
	// Get latest featured articles
	articles, err := s.articleRepo.Article.GetAllWithRelations()
	if err != nil {
		return "", fmt.Errorf("failed to get articles: %w", err)
	}

	// Filter for featured articles and limit
	var featuredArticles []db.Article
	for _, article := range articles {
		if article.IsFeatured && len(featuredArticles) < limit {
			featuredArticles = append(featuredArticles, article)
		}
	}

	// Prepare template data
	templateData := map[string]interface{}{
		"Articles":       featuredArticles,
		"Date":           time.Now().Format("January 2, 2006"),
		"Year":           time.Now().Year(),
		"SubscriberName": "Preview User",
	}

	// Merge custom data if provided
	for key, value := range customData {
		templateData[key] = value
	}

	// Render template
	return s.emailService.RenderTemplate("newsletter", templateData)
}

// GetNewsletterTemplate returns the current newsletter template content
func (s *NewsletterService) GetNewsletterTemplate() (string, error) {
	return s.emailService.GetTemplateContent("newsletter")
}

// UpdateNewsletterTemplate updates the newsletter template
func (s *NewsletterService) UpdateNewsletterTemplate(templateContent string) error {
	return s.emailService.SaveTemplateContent("newsletter", templateContent)
}
