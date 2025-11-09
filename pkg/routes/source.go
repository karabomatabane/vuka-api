package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterSourceRoutes = func(router *mux.Router) {
	sourceController := controllers.NewSourceController()

	// Public routes (no authentication required)
	sourceRouter := router.PathPrefix("/source").Subrouter()
	sourceRouter.HandleFunc("", sourceController.CreateSource).Methods(http.MethodPost)
	sourceRouter.HandleFunc("", sourceController.GetAllSources)
	sourceRouter.HandleFunc("/{id}", sourceController.GetSourceByID).Methods(http.MethodGet)

	// Protected routes (authentication required)
	protectedRouter := sourceRouter.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyTokenAndAdmin)
	// Admin-only routes for source management
	protectedRouter.HandleFunc("/{id}", sourceController.UpdateSource).Methods(http.MethodPatch)
	protectedRouter.HandleFunc("/{id}", sourceController.DeleteSource).Methods(http.MethodDelete)
	protectedRouter.HandleFunc("/{id}/ingest", sourceController.IngestSourceFeed).Methods(http.MethodPost)
}
