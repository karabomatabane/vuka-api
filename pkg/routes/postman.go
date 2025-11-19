package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"

	"github.com/gorilla/mux"
)

// RegisterPostmanRoutes registers routes for Postman collection generation
func RegisterPostmanRoutes(router *mux.Router) {
	postmanController := controllers.NewPostmanController(router)

	// Public endpoint to download Postman collection
	router.HandleFunc("/api/postman/collection", postmanController.GenerateCollection).
		Methods(http.MethodGet)
}
