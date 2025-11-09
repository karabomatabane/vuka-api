package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterDirectoryRoutes = func(router *mux.Router) {
	controller := controllers.NewDirectoryController()
	directoryRouter := router.PathPrefix("/directory").Subrouter()
	// Public routes
	directoryRouter.HandleFunc("", controller.GetAllDirectories).Methods("GET")
	directoryRouter.HandleFunc("/entries/{category_id}", controller.GetDirectoryEntriesByCategoryID).Methods("GET")

	// Protected routes (authentication required)
	protectedRouter := directoryRouter.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyTokenAndAdmin)
	protectedRouter.Handle("/overview", middleware.VerifyToken(http.HandlerFunc(controller.GetDirectoryOverview))).Methods("GET")
	protectedRouter.HandleFunc("", controller.CreateDirectoryCategory).Methods("POST")
	protectedRouter.HandleFunc("/entries", controller.CreateDirectoryEntry).Methods("POST")
	protectedRouter.HandleFunc("/entry/{entry_id}", controller.GetDirectoryEntryByID).Methods("GET")
	// router.HandleFunc("/directory/{id}", controller.UpdateDirectory).Methods("PUT")
	// router.HandleFunc("/directory/{id}", controller.DeleteDirectory).Methods("DELETE")
}