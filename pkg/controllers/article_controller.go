package controllers

import (
	"net/http"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

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

	article, err := fc.articleService.GetArticleByID(vars["id"])
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
	_, err := fc.articleService.GetArticleByID(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Article ID does not exist", http.StatusBadRequest)
		return
	}

	updates := make(map[string]any)
	if err := httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedArticle, err := fc.articleService.UpdateArticle(vars["id"], updates)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, updatedArticle)
}

func (fc *ArticleController) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := fc.articleService.DeleteArticle(vars["id"]); err != nil {
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
