package routes

import (
	"vuka-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterCategoryRoutes = func(router *mux.Router) {
	controller := controllers.NewCategoryController()
	router.HandleFunc("/category", controller.GetAllCategories).Methods("GET")
}