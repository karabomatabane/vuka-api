package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/middleware"
)

var RegisterUserRoutes = func(router *mux.Router) {
	userController := controllers.NewUserController()

	// Protected routes (authentication required)
	protectedRouter := router.PathPrefix("/").Subrouter()
	protectedRouter.Use(func(next http.Handler) http.Handler {
		return middleware.VerifyTokenAndAdmin(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	// Admin-only routes for user management
	protectedRouter.HandleFunc("/user",
		userController.GetAllUsers).
		Methods(httpx.HTTP_GET)
}
