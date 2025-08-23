package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
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
