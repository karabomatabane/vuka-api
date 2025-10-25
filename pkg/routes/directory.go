package routes

import (
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"

	"github.com/gorilla/mux"
)

var RegisterDirectoryRoutes = func(router *mux.Router) {
	controller := controllers.NewDirectoryController()

	// Public routes
	router.HandleFunc("/directory", controller.GetAllDirectories).Methods("GET")
	router.HandleFunc("/directory/entries/{category_id}", controller.GetDirectoryEntriesByCategoryID).Methods("GET")

	// Protected routes
	router.Handle("/directory/overview", middleware.VerifyToken(http.HandlerFunc(controller.GetDirectoryOverview))).Methods("GET")
	router.HandleFunc("/directory", controller.CreateDirectoryCategory).Methods("POST")
	router.HandleFunc("/directory/entries", controller.CreateDirectoryEntry).Methods("POST")
	// router.HandleFunc("/directory/{id}", controller.GetDirectoryByID).Methods("GET")
	// router.HandleFunc("/directory/{id}", controller.UpdateDirectory).Methods("PUT")
	// router.HandleFunc("/directory/{id}", controller.DeleteDirectory).Methods("DELETE")
}