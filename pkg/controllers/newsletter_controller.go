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
