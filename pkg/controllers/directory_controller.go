package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"
)

// DirectoryController is a controller for directory-related operations.
type DirectoryController struct {
	service *services.DirectoryService
}

// NewDirectoryController creates a new DirectoryController.
func NewDirectoryController() *DirectoryController {
	serviceManager := services.NewServices(config.GetDB())
	return &DirectoryController{service: serviceManager.Directory}
}

// GetAllCategories handles the request to get all categories.
func (c *DirectoryController) GetAllDirectories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	directories, err := c.service.GetAllDirectories()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	httpx.WriteJSON(w, http.StatusOK, directories)
}

// GetDirectoryEntriesByCategoryID handles the request to get all directory entries for a specific category.
func (c *DirectoryController) GetDirectoryEntriesByCategoryID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryIDStr := r.URL.Query().Get("category_id")
	if categoryIDStr == "" {
		httpx.WriteErrorJSON(w, "category_id is required", http.StatusBadRequest)
	}

	entries, err := c.service.GetDirectoryEntriesByCategoryID(categoryIDStr)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	httpx.WriteJSON(w, http.StatusOK, entries)
}

// CreateDirectoryCategory handles the request to create a new directory category.
func (c *DirectoryController) CreateDirectoryCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category db.DirectoryCategory
	if err := httpx.ParseBody(r, &category); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateDirectoryCategory(&category); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, category)
}

// CreateDirectoryEntry handles the request to create a new directory entry.
func (c *DirectoryController) CreateDirectoryEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var entry db.DirectoryEntry
	if err := httpx.ParseBody(r, &entry); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateDirectoryEntry(&entry); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, entry)
}
