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
}
