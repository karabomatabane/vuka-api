package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	userController := controllers.NewUserController()

	// Protected routes (authentication required)
	protectedRouter := router.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyTokenAndAdmin)
	// Admin-only routes for user management
	protectedRouter.HandleFunc("/user",
		userController.GetAllUsers).
		Methods(http.MethodGet)
	protectedRouter.HandleFunc("/user/{id}",
		userController.GetUserByID).
		Methods(http.MethodGet)
	protectedRouter.HandleFunc("/user/{id}",
		userController.UpdateUser).
		Methods(http.MethodPatch)
	protectedRouter.HandleFunc("/user/{id}",
		userController.DeleteUser).
		Methods(http.MethodDelete)
	protectedRouter.HandleFunc("/user/{id}/role",
		userController.UpdateUserRole).
		Methods(http.MethodPatch)
}
