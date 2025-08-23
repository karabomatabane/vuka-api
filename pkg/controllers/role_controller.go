package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"
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
