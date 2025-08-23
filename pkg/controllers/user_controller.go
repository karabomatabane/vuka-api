package controllers

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/user"
	"vuka-api/pkg/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	serviceManager := services.NewServices(config.GetDB())
	return &UserController{
		userService: serviceManager.User,
	}
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSON(w, http.StatusOK, users)
}

func (uc *UserController) UpdateUserRole(w http.ResponseWriter, r *http.Request) {
	var body user.UpdateUserRoleBody
	if err := httpx.ParseBody(r, body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := uc.userService.UpdateUserRole(body)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, updatedUser)
}
