package controllers

import (
	"net/http"
	"time"
	"vuka-api/pkg/config"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type FilmController struct {
	filmService *services.FilmService
}

func NewFilmController() *FilmController {
	serviceManager := services.NewServices(config.GetDB())
	return &FilmController{
		filmService: serviceManager.Film,
	}
}

func (fc *FilmController) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	filmID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	film, err := fc.filmService.GetFilmByID(filmID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, film)
}

func (fc *FilmController) GetAllFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := fc.filmService.GetAllFilms()
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, films)
}

func (fc *FilmController) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	filmID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	// Get existing film
	existingFilm, err := fc.filmService.GetFilmByID(filmID)
	if err != nil {
		httpx.WriteErrorJSON(w, "Film not found", http.StatusNotFound)
		return
	}

	// Parse request body for updates
	var updates db.Film
	if err = httpx.ParseBody(r, &updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	updates.ID = existingFilm.ID
	if err = fc.filmService.UpdateFilm(&updates); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return updated film
	updatedFilm, err := fc.filmService.GetFilmByID(filmID)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, updatedFilm)
}

func (fc *FilmController) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	filmID, err := uuid.Parse(vars["id"])
	if err != nil {
		httpx.WriteErrorJSON(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	if err = fc.filmService.DeleteFilm(filmID); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (fc *FilmController) CreateFilm(w http.ResponseWriter, r *http.Request) {
	var film db.Film
	if err := httpx.ParseBody(r, &film); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := fc.filmService.CreateFilm(&film); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, film)
}

func (fc *FilmController) CreateFilmFromTmdb(w http.ResponseWriter, r *http.Request) {
	type TmdbFilmBody struct {
		ScreeningDate time.Time `json:"screeningDate"`
	}

	vars := mux.Vars(r)
	var body TmdbFilmBody
	if err := httpx.ParseBody(r, &body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	film, err := fc.filmService.CreateFromTmdb(vars["id"], body.ScreeningDate)
	if err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, film)
}

func (fc *FilmController) ImportDirectorImages(w http.ResponseWriter, r *http.Request) {
	type DirectorBody struct {
		Name   string `json:"name,omitempty"`
		ImgUrl string `json:"ImgUrl,omitempty"`
	}

	var body []DirectorBody
	if err := httpx.ParseBody(r, &body); err != nil {
		httpx.WriteErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, directorBody := range body {
		err := fc.filmService.UpdateDirectorImage(directorBody.Name, directorBody.ImgUrl)
		if err != nil {
			httpx.WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}
