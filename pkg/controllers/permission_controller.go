package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/gorilla/mux"
)

type PermissionController struct {
	permissionService *services.PermissionService
}

func NewPermissionController() *PermissionController {
	serviceManager := services.NewServices(config.GetDB())
	return &PermissionController{
		permissionService: serviceManager.Permission,
	}
}

// Permission CRUD
func (pc *PermissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission db.Permission
	if err := httpx.ParseBody(r, &permission); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pc.permissionService.CreatePermission(&permission); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, permission)
}

func (pc *PermissionController) GetAllPermissions(w http.ResponseWriter, _ *http.Request) {
	permissions, err := pc.permissionService.GetAllPermissions()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, permissions)
}

func (pc *PermissionController) GetPermissionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	permission, err := pc.permissionService.GetPermissionByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Permission not found", http.StatusNotFound)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, permission)
}

func (pc *PermissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	permission, err := pc.permissionService.GetPermissionByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Permission not found", http.StatusNotFound)
		return
	}

	var updates map[string]any
	if err := httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if name, ok := updates["name"].(string); ok {
		permission.Name = name
	}

	if err := pc.permissionService.UpdatePermission(permission); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, permission)
}

func (pc *PermissionController) DeletePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := pc.permissionService.DeletePermission(vars["id"]); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Section CRUD
func (pc *PermissionController) CreateSection(w http.ResponseWriter, r *http.Request) {
	var section db.Section
	if err := httpx.ParseBody(r, &section); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := pc.permissionService.CreateSection(&section); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, section)
}

func (pc *PermissionController) GetAllSections(w http.ResponseWriter, _ *http.Request) {
	sections, err := pc.permissionService.GetAllSections()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, sections)
}

func (pc *PermissionController) GetSectionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section, err := pc.permissionService.GetSectionByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Section not found", http.StatusNotFound)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, section)
}

func (pc *PermissionController) UpdateSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	section, err := pc.permissionService.GetSectionByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Section not found", http.StatusNotFound)
		return
	}

	var updates map[string]any
	if err := httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if name, ok := updates["name"].(string); ok {
		section.Name = name
	}

	if err := pc.permissionService.UpdateSection(section); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, section)
}

func (pc *PermissionController) DeleteSection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := pc.permissionService.DeleteSection(vars["id"]); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
