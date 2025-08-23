package routes

import (
	"github.com/gorilla/mux"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/middleware"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	authController := controllers.NewAuthController()
	router.HandleFunc("/auth/register",
		authController.Register).
		Methods(httpx.HTTP_POST)
	router.HandleFunc("/auth/login",
		authController.Login).
		Methods(httpx.HTTP_POST)

	// Protected routes (authentication required)
	protectedRouter := router.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyToken)

	// Admin-only routes for account code management
	protectedRouter.HandleFunc("/auth/account-code", authController.CreateAccountCode).Methods(httpx.HTTP_POST)
	protectedRouter.HandleFunc("/auth/account-code", authController.GetAccountCode).Methods(httpx.HTTP_GET)
	router.HandleFunc("/auth/account-code", authController.DeleteAccountCode).Methods(httpx.HTTP_DELETE)
}
