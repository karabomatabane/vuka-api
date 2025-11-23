package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type NewsletterController struct {
	newsletterService *services.NewsletterService
}

func NewNewsletterController() *NewsletterController {
	serviceManager := services.NewServices(config.GetDB())
	return &NewsletterController{
		newsletterService: serviceManager.Newsletter,
	}
}

func (nc *NewsletterController) Subscribe(w http.ResponseWriter, r *http.Request) {
	var subscriber db.NewsletterSubscriber
	if err := httpx.ParseBody(r, &subscriber); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(subscriber); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nc.newsletterService.Subscribe(&subscriber); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, subscriber)
}

func (nc *NewsletterController) GetAllSubscribers(w http.ResponseWriter, _ *http.Request) {
	subscribers, err := nc.newsletterService.GetAllSubscribers()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, subscribers)
}

func (nc *NewsletterController) GetSubscriberByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	subscriber, err := nc.newsletterService.GetSubscriberByID(id)
	if err != nil {
		httpx.WriteErrorJSON(w, "Subscriber not found", http.StatusNotFound)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, subscriber)
}

func (nc *NewsletterController) UpdateSubscriber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	existingSubscriber, err := nc.newsletterService.GetSubscriberByID(id)
	if err != nil {
		httpx.WriteErrorJSON(w, "Subscriber not found", http.StatusNotFound)
		return
	}

	var updates map[string]any
	if err := httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if preferredName, ok := updates["preferredName"].(string); ok {
		existingSubscriber.PreferredName = preferredName
	}
	if email, ok := updates["email"].(string); ok {
		existingSubscriber.Email = email
	}
	if phoneNumber, ok := updates["phoneNumber"].(string); ok {
		existingSubscriber.PhoneNumber = phoneNumber
	}

	if err := nc.newsletterService.UpdateSubscriber(existingSubscriber); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, existingSubscriber)
}

func (nc *NewsletterController) DeleteSubscriber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := nc.newsletterService.DeleteSubscriber(id); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SendNewsletter sends a newsletter to all subscribers
func (nc *NewsletterController) SendNewsletter(w http.ResponseWriter, r *http.Request) {
	type NewsletterRequest struct {
		Subject      string                 `json:"subject" validate:"required"`
		Content      string                 `json:"content"`
		UseTemplate  bool                   `json:"useTemplate"`
		TemplateData map[string]interface{} `json:"templateData"`
	}

	var req NewsletterRequest
	if err := httpx.ParseBody(r, &req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nc.newsletterService.SendNewsletter(req.Subject, req.Content, req.UseTemplate, req.TemplateData); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Newsletter sent successfully",
	})
}

// SendNewsletterWithArticles sends a newsletter with latest featured articles
func (nc *NewsletterController) SendNewsletterWithArticles(w http.ResponseWriter, r *http.Request) {
	type ArticleNewsletterRequest struct {
		Subject string `json:"subject" validate:"required"`
		Limit   int    `json:"limit"`
	}

	var req ArticleNewsletterRequest
	if err := httpx.ParseBody(r, &req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Limit == 0 {
		req.Limit = 5 // Default to 5 articles
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nc.newsletterService.SendNewsletterWithLatestArticles(req.Subject, req.Limit); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Newsletter with articles sent successfully",
	})
}

// SendTestEmail sends a test email
func (nc *NewsletterController) SendTestEmail(w http.ResponseWriter, r *http.Request) {
	type TestEmailRequest struct {
		Email string `json:"email" validate:"required,email"`
		Name  string `json:"name" validate:"required"`
	}

	var req TestEmailRequest
	if err := httpx.ParseBody(r, &req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nc.newsletterService.SendTestEmail(req.Email, req.Name); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Test email sent successfully",
	})
}

// TestSMTPConnection tests the SMTP connection
func (nc *NewsletterController) TestSMTPConnection(w http.ResponseWriter, _ *http.Request) {
	if err := nc.newsletterService.TestSMTPConnection(); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "SMTP connection successful",
	})
}

// PreviewNewsletter generates a preview of the newsletter email
func (nc *NewsletterController) PreviewNewsletter(w http.ResponseWriter, r *http.Request) {
	type PreviewRequest struct {
		ArticleLimit int                    `json:"articleLimit"`
		TemplateData map[string]interface{} `json:"templateData"`
	}

	var req PreviewRequest
	if err := httpx.ParseBody(r, &req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.ArticleLimit == 0 {
		req.ArticleLimit = 5
	}

	htmlContent, err := nc.newsletterService.GenerateNewsletterPreview(req.ArticleLimit, req.TemplateData)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlContent))
}

// GetTemplate returns the current newsletter template
func (nc *NewsletterController) GetTemplate(w http.ResponseWriter, _ *http.Request) {
	template, err := nc.newsletterService.GetNewsletterTemplate()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"template": template,
	})
}

// UpdateTemplate updates the newsletter template
func (nc *NewsletterController) UpdateTemplate(w http.ResponseWriter, r *http.Request) {
	type TemplateRequest struct {
		Template string `json:"template" validate:"required"`
	}

	var req TemplateRequest
	if err := httpx.ParseBody(r, &req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := nc.newsletterService.UpdateNewsletterTemplate(req.Template); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Template updated successfully",
	})
}
