package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterNewsletterRoutes = func(router *mux.Router) {
	newsletterController := controllers.NewNewsletterController()

	// Public route - anyone can subscribe
	router.HandleFunc("/newsletter/subscribe",
		newsletterController.Subscribe).
		Methods(http.MethodPost)

	// Protected routes (authentication required)
	protectedRouter := router.PathPrefix("/newsletter").Subrouter()
	protectedRouter.Use(middleware.VerifyTokenAndAdmin)

	protectedRouter.HandleFunc("/subscribers",
		newsletterController.GetAllSubscribers).
		Methods(http.MethodGet)

	protectedRouter.HandleFunc("/subscribers/{id}",
		newsletterController.GetSubscriberByID).
		Methods(http.MethodGet)

	protectedRouter.HandleFunc("/subscribers/{id}",
		newsletterController.UpdateSubscriber).
		Methods(http.MethodPatch)

	protectedRouter.HandleFunc("/subscribers/{id}",
		newsletterController.DeleteSubscriber).
		Methods(http.MethodDelete)

	// Email sending routes
	protectedRouter.HandleFunc("/send",
		newsletterController.SendNewsletter).
		Methods(http.MethodPost)

	protectedRouter.HandleFunc("/send/articles",
		newsletterController.SendNewsletterWithArticles).
		Methods(http.MethodPost)

	protectedRouter.HandleFunc("/test-email",
		newsletterController.SendTestEmail).
		Methods(http.MethodPost)

	protectedRouter.HandleFunc("/test-smtp",
		newsletterController.TestSMTPConnection).
		Methods(http.MethodGet)

	// Template management routes
	protectedRouter.HandleFunc("/template",
		newsletterController.GetTemplate).
		Methods(http.MethodGet)

	protectedRouter.HandleFunc("/template",
		newsletterController.UpdateTemplate).
		Methods(http.MethodPut)

	protectedRouter.HandleFunc("/preview",
		newsletterController.PreviewNewsletter).
		Methods(http.MethodPost)
}
