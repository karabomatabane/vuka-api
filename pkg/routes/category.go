package routes

import (
	"vuka-api/pkg/controllers"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(router *mux.Router, controller *controllers.CategoryController) {
	router.HandleFunc("/category", controller.GetAllCategories).Methods("GET")
}