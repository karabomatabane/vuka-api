package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterPermissionRoutes = func(router *mux.Router) {
	permissionController := controllers.NewPermissionController()

	// Permission routes
	permissionRouter := router.PathPrefix("/permission").Subrouter()
	
	// Public permission routes
	permissionRouter.HandleFunc("", permissionController.GetAllPermissions).
		Methods(http.MethodGet)
	permissionRouter.HandleFunc("/{id}", permissionController.GetPermissionByID).
		Methods(http.MethodGet)

	// Protected permission routes (admin only)
	permissionRouter.HandleFunc("", middleware.VerifyTokenAndAdminFunc(permissionController.CreatePermission)).
		Methods(http.MethodPost)
	permissionRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(permissionController.UpdatePermission)).
		Methods(http.MethodPatch)
	permissionRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(permissionController.DeletePermission)).
		Methods(http.MethodDelete)

	// Section routes
	sectionRouter := router.PathPrefix("/section").Subrouter()
	
	// Public section routes
	sectionRouter.HandleFunc("", permissionController.GetAllSections).
		Methods(http.MethodGet)
	sectionRouter.HandleFunc("/{id}", permissionController.GetSectionByID).
		Methods(http.MethodGet)

	// Protected section routes (admin only)
	sectionRouter.HandleFunc("", middleware.VerifyTokenAndAdminFunc(permissionController.CreateSection)).
		Methods(http.MethodPost)
	sectionRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(permissionController.UpdateSection)).
		Methods(http.MethodPatch)
	sectionRouter.HandleFunc("/{id}", middleware.VerifyTokenAndAdminFunc(permissionController.DeleteSection)).
		Methods(http.MethodDelete)
}
