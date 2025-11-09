package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/user"
	"vuka-api/pkg/services"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	existingUser, err := uc.userService.GetUserByID(userID)
	if err != nil {
		httpx.WriteErrorJSON(w, "User not found", http.StatusNotFound)
		return
	}

	var updates map[string]any
	if err := httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if username, ok := updates["username"].(string); ok {
		existingUser.Username = username
	}
	if roleID, ok := updates["roleId"].(string); ok {
		roleUUID, err := uuid.Parse(roleID)
		if err != nil {
			httpx.WriteErrorJSON(w, "Invalid role ID", http.StatusBadRequest)
			return
		}
		existingUser.RoleID = roleUUID
	}

	if err := uc.userService.UpdateUser(existingUser); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, existingUser)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if err := uc.userService.DeleteUser(userID); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := uc.userService.GetUserByID(userID)
	if err != nil {
		httpx.WriteErrorJSON(w, "User not found", http.StatusNotFound)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, user)
}
