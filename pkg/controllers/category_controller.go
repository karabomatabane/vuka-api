package controllers

import (
	"net/http"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/services"
)

// CategoryController is a controller for category-related operations.
type CategoryController struct {
	service *services.CategoryService
}

// NewCategoryController creates a new CategoryController.
func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// GetAllCategories handles the request to get all categories.
func (c *CategoryController) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := c.service.GetAllCategories()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, categories)
}