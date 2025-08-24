package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/gorilla/mux"
)

type RoleController struct {
	roleService *services.RoleService
}

func NewRoleController() *RoleController {
	serviceManager := services.NewServices(config.GetDB())
	return &RoleController{
		roleService: serviceManager.Role,
	}
}

func (c *RoleController) Create(w http.ResponseWriter, r *http.Request) {
	var role db.Role
	if err := httpx.ParseBody(r, &role); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.roleService.Create(&role); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, role)
}

func (c *RoleController) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := c.roleService.GetAll()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, roles)
}

func (c *RoleController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role, err := c.roleService.GetById(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, role)
}

func (c *RoleController) GetWithPermissions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role, err := c.roleService.GetWithPermissions(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, role)
}

func (c *RoleController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role, err := c.roleService.GetById(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Role ID does not exist", http.StatusInternalServerError)
		return
	}

	var updates map[string]any
	err = httpx.ParseBody(r, &updates)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name, ok := updates["name"].(string)
	if !ok {
		httpx.WriteErrorJSON(w, "Role name not provided", http.StatusInternalServerError)
		return
	}
	role.Name = name
	err = c.roleService.Update(role)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, role)
}

func (c *RoleController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := c.roleService.Delete(vars["id"]); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}