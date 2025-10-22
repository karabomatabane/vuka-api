package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/services"
)

// CategoryController is a controller for category-related operations.
type CategoryController struct {
	service *services.CategoryService
}

// NewCategoryController creates a new CategoryController.
func NewCategoryController() *CategoryController {
	serviceManager := services.NewServices(config.GetDB())
	return &CategoryController{service: serviceManager.Category}
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