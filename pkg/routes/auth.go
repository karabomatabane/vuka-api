package routes

import (
	"github.com/gorilla/mux"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/httpx"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	authController := controllers.NewAuthController()
	router.HandleFunc("/auth/register",
		authController.Register).
		Methods(httpx.HTTP_POST)
	router.HandleFunc("/auth/login",
		authController.Login).
		Methods(httpx.HTTP_POST)
}
