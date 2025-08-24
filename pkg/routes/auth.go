package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	authController := controllers.NewAuthController()
	router.HandleFunc("/auth/register",
		authController.Register).
		Methods(http.MethodPost)
	router.HandleFunc("/auth/login",
		authController.Login).
		Methods(http.MethodPost)
}
