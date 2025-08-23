package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"
)

var RegisterRoleRoutes = func(router *mux.Router) {
	roleController := controllers.NewRoleController()

	// Public routes (no authentication required)
	roleRouter := router.PathPrefix("/role").Subrouter()
	roleRouter.HandleFunc("", roleController.Create).
		Methods(http.MethodPost)

	// Protected routes (authentication required)
	protectedRouter := roleRouter.PathPrefix("/").Subrouter()
	protectedRouter.Use(func(next http.Handler) http.Handler {
		return middleware.VerifyTokenAndAdmin(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

	// Admin-only routes for role management

}
