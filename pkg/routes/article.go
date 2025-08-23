package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"vuka-api/pkg/controllers"
	"vuka-api/pkg/middleware"
)

var RegisterArticleRoutes = func(router *mux.Router) {
	articleController := controllers.NewArticleController()

	// Public routes (no authentication required)
	articleRouter := router.PathPrefix("/article").Subrouter()
	articleRouter.HandleFunc("", articleController.GetAllArticles).
		Methods(http.MethodGet)
	articleRouter.HandleFunc("/{id}", articleController.GetArticle).
		Methods(http.MethodGet)

	// Protected routes (authentication required)
	protectedRouter := articleRouter.PathPrefix("/").Subrouter()
	protectedRouter.Use(middleware.VerifyToken)

	// Admin-only routes for article management
	protectedRouter.HandleFunc("", articleController.CreateArticle).
		Methods(http.MethodPost)
	protectedRouter.HandleFunc("/{id}", articleController.UpdateArticle).
		Methods(http.MethodPut)
	protectedRouter.HandleFunc("/{id}", articleController.DeleteArticle).
		Methods(http.MethodDelete)
}
