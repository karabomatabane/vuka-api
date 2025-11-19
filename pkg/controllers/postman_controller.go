package controllers

import (
	"encoding/json"
	"net/http"
	"vuka-api/pkg/postman"

	"github.com/gorilla/mux"
)

// PostmanController handles Postman collection generation
type PostmanController struct {
	router *mux.Router
}

// NewPostmanController creates a new Postman controller
func NewPostmanController(router *mux.Router) *PostmanController {
	return &PostmanController{
		router: router,
	}
}

// GenerateCollection generates a Postman collection and returns it as JSON
func (pc *PostmanController) GenerateCollection(w http.ResponseWriter, r *http.Request) {
	baseURL := r.URL.Query().Get("baseUrl")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	generator := postman.NewGenerator(baseURL, "Vuka API")
	generator.Description = "Complete API documentation for Vuka API"

	collection, err := generator.Generate(pc.router)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to generate collection: " + err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=vuka-api.postman_collection.json")

	if err := json.NewEncoder(w).Encode(collection); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Failed to encode collection: " + err.Error(),
		})
	}
}
