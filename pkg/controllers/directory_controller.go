package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/middleware"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

// GetDirectoryOverview handles the request to get an overview of the directory.
func (c *DirectoryController) GetDirectoryOverview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, ok := r.Context().Value(middleware.UserContextKey).(jwt.MapClaims)
	if !ok {
		httpx.WriteErrorJSON(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := uuid.Parse(user["userId"].(string))
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	response, err := c.service.GetDirectoryOverview(userID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, response)
}

// GetAllCategories handles the request to get all categories.
func (c *DirectoryController) GetAllDirectories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	directories, err := c.service.GetAllDirectories()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, directories)
}

// GetDirectoryEntriesByCategoryID handles the request to get a directory category with all its entries.
func (c *DirectoryController) GetDirectoryEntriesByCategoryID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryID := mux.Vars(r)["category_id"]
	if categoryID == "" {
		httpx.WriteErrorJSON(w, "category_id is required", http.StatusBadRequest)
		return
	}

	category, err := c.service.GetDirectoryEntriesByCategoryID(categoryID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, category)
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

	// Validate contact info types before creation
	for i := range entry.ContactInfo {
		if !entry.ContactInfo[i].Type.IsValid() {
			httpx.WriteErrorJSON(w, "Invalid contact type: "+string(entry.ContactInfo[i].Type), http.StatusBadRequest)
			return
		}
		// Ensure the DirectoryEntryID is not set (it will be set automatically)
		entry.ContactInfo[i].DirectoryEntryID = uuid.Nil
	}

	if err := c.service.CreateDirectoryEntry(&entry); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, entry)
}

// GetDirectoryEntryByID handles the request for getting an entry by Id.
func (c *DirectoryController) GetDirectoryEntryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	entryID := mux.Vars(r)["entry_id"]

	if entryID == "" {
		httpx.WriteErrorJSON(w, "entry_id is required", http.StatusBadRequest)
		return
	}

	entry, err := c.service.GetDirectoryEntryByID(entryID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	httpx.WriteJSON(w, http.StatusOK, entry)
}
