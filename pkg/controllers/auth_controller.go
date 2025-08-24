package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models"
	"vuka-api/pkg/services"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController() *AuthController {
	serviceManager := services.NewServices(config.GetDB())
	return &AuthController{
		authService: serviceManager.Auth,
	}
}

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterBody
	if err := httpx.ParseBody(r, &body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if body.ConfirmPassword != body.Password {
		httpx.WriteErrorJSON(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	authResponse, err := ac.authService.Register(body)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSON(w, http.StatusCreated, authResponse)
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	body := models.LoginBody{}
	if err := httpx.ParseBody(r, &body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := ac.authService.Login(body)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSON(w, http.StatusOK, authResponse)
}
