package routes

import (
	"github.com/gorilla/mux"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/middleware"
)

var RegisterFilmRoutes = func(router *mux.Router) {
	// Initialize film controller
	filmController := controllers.NewFilmController()

	// Public routes (no authentication required)
	router.HandleFunc("/film", filmController.GetAllFilms).
		Methods(httpx.HTTP_GET)
	router.HandleFunc("/film/{id}", filmController.GetFilm).
		Methods(httpx.HTTP_GET)
	router.HandleFunc("/admin/directors/import/batch", filmController.ImportDirectorImages).
		Methods(httpx.HTTP_POST)

	// Protected routes (authentication required)
	protectedRouter := router.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyToken)

	// Admin-only routes for film management
	protectedRouter.HandleFunc("/film", filmController.CreateFilm).
		Methods(httpx.HTTP_POST)
	protectedRouter.HandleFunc("/film/tmdb/{id}", filmController.CreateFilmFromTmdb).
		Methods(httpx.HTTP_POST)
	protectedRouter.HandleFunc("/film/{id}", filmController.UpdateFilm).
		Methods(httpx.HTTP_PUT)
	protectedRouter.HandleFunc("/film/{id}", filmController.DeleteFilm).
		Methods(httpx.HTTP_DELETE)
}
