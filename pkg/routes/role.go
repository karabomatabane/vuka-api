package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterRoleRoutes = func(router *mux.Router) {
	roleController := controllers.NewRoleController()

	roleRouter := router.PathPrefix("/role").Subrouter()

	// Public routes (no authentication required)
	roleRouter.HandleFunc("", roleController.GetAll).
		Methods(http.MethodGet)
	roleRouter.HandleFunc("/{id}", roleController.GetById).
		Methods(http.MethodGet)
	roleRouter.HandleFunc("/{id}/permissions", roleController.GetWithPermissions).
		Methods(http.MethodGet)

	// Protected routes (admin only)
	roleRouter.HandleFunc("", middleware.VerifyTokenAndAdminFunc(roleController.Create)).
		Methods(http.MethodPost)
	roleRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(roleController.Update)).
		Methods(http.MethodPatch)
	roleRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(roleController.Delete)).
		Methods(http.MethodDelete)

	// Role permission management (admin only)
	roleRouter.HandleFunc("/permissions", middleware.VerifyTokenAndAdminFunc(roleController.AssignPermissionToRole)).
		Methods(http.MethodPost)
	roleRouter.HandleFunc("/{roleId}/permissions/{sectionId}/{permissionId}",
		middleware.VerifyTokenAndAdminFunc(roleController.RemovePermissionFromRole)).
		Methods(http.MethodDelete)
}
