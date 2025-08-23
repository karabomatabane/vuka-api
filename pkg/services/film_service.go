package services

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/models/mappers"
	"vuka-api/pkg/repository"
)

type FilmService struct {
	repos *repository.Repositories
}

func NewFilmService(repos *repository.Repositories) *FilmService {
	return &FilmService{repos: repos}
}

func (s *FilmService) CreateFilm(film *db.Film) error {
	return s.repos.Film.Create(film)
}

func (s *FilmService) CreateFromTmdb(tmdbFilmId string, screeningDate time.Time) (*db.Film, error) {
	// Fetch film details
	filmDetailURL := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?language=en-US'", tmdbFilmId)
	tmdbResp, err := httpx.GetFromTmdb(filmDetailURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch film details: %w", err)
	}
	defer tmdbResp.Body.Close()

	// Fetch cast details
	filmCastURL := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s/credits?language=en-US'", tmdbFilmId)
	tmdbCastResp, err := httpx.GetFromTmdb(filmCastURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cast details: %w", err)
	}
	defer tmdbCastResp.Body.Close()

	// Check HTTP status codes
	if tmdbResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("film details request failed: %s", tmdbResp.Status)
	}
	if tmdbCastResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("film cast request failed: %s", tmdbCastResp.Status)
	}

	// After checking HTTP status codes
	filmData, err := io.ReadAll(tmdbResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read film details: %w", err)
	}
	castData, err := io.ReadAll(tmdbCastResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read cast details: %w", err)
	}

	film, _, _, err := mappers.MapTmdbToFilm(filmData, castData, screeningDate)
	if err != nil {
		return nil, fmt.Errorf("failed to map TMDb data: %w", err)
	}

	genres := make([]db.Genre, len(film.Genres))
	for i, g := range film.Genres {
		genre, err := s.repos.Genre.FindOrCreate(g.Name)
		if err != nil {
			return nil, err
		}
		// pass value to array
		genres[i] = *genre
	}

	// update model with genres from db
	film.Genres = genres

	// Save film to DB
	if err := s.repos.Film.Create(film); err != nil {
		return nil, err
	}

	return film, nil
}

func (s *FilmService) GetFilmByID(id uuid.UUID) (*db.Film, error) {
	return s.repos.Film.GetWithRelations(id)
}

func (s *FilmService) GetAllFilms() ([]db.Film, error) {
	return s.repos.Film.GetAll()
}

func (s *FilmService) UpdateFilm(film *db.Film) error {
	return s.repos.Film.Update(film)
}

func (s *FilmService) DeleteFilm(id uuid.UUID) error {
	return s.repos.Film.Delete(id)
}

func (s *FilmService) UpdateDirectorImage(name, imgUrl string) error {
	dir, err := s.repos.Director.Update(name, imgUrl)
	if err != nil {
		return err
	}
	fmt.Printf("Updated: %s:%s", dir.Name, dir.ImgUrl)
	return nil
}
