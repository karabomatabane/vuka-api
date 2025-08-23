package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	articleService *services.ArticleService
}

func NewArticleController() *ArticleController {
	serviceManager := services.NewServices(config.GetDB())
	return &ArticleController{
		articleService: serviceManager.Article,
	}
}

func (fc *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := fc.articleService.GetArticleByID(articleID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, article)
}

func (fc *ArticleController) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	articles, err := fc.articleService.GetAllArticles()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, articles)
}

func (fc *ArticleController) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	// Get existing article
	existingArticle, err := fc.articleService.GetArticleByID(articleID)
	if err != nil {
		httpx.WriteErrorJSON(w, "Article not found", http.StatusNotFound)
		return
	}

	// Parse request body for updates
	var updates db.Article
	if err = httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	updates.ID = existingArticle.ID
	if err = fc.articleService.UpdateArticle(&updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated article
	updatedArticle, err := fc.articleService.GetArticleByID(articleID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, updatedArticle)
}

func (fc *ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	articleID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	if err = fc.articleService.DeleteArticle(articleID); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (fc *ArticleController) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article *db.Article
	if err := httpx.ParseBody(r, article); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := fc.articleService.CreateArticle(article); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, article)
}
