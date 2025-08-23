package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models"
	"vuka-api/pkg/services"
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
	if err := httpx.ParseBody(r, body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
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
	if err := httpx.ParseBody(r, body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	authResponse, err := ac.authService.Login(body)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSON(w, http.StatusOK, authResponse)
}

func (ac *AuthController) CreateAccountCode(w http.ResponseWriter, r *http.Request) {
	var body models.AccountCode
	if err := httpx.ParseBody(r, body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountCode, err := ac.authService.CreateAccountCode(body)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}

	httpx.WriteJSON(w, http.StatusCreated, accountCode)
}

func (ac *AuthController) GetAccountCode(w http.ResponseWriter, _ *http.Request) {
	accountCode, err := ac.authService.GetAccountCode()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, accountCode)
}

func (ac *AuthController) DeleteAccountCode(w http.ResponseWriter, _ *http.Request) {
	err := ac.authService.DeleteAccountCode()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
	httpx.WriteJSON(w, http.StatusOK, ac)
}
