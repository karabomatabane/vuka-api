package mappers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"vuka-api/pkg/models/db"
)

type TmdbFilm struct {
	Title            string                  `json:"title"`
	PosterPath       string                  `json:"poster_path"`
	ReleaseDate      string                  `json:"release_date"`
	Runtime          int                     `json:"runtime"`
	Genres           []struct{ Name string } `json:"genres"`
	SpokenLanguages  []struct{ Name string } `json:"spoken_languages"`
	Overview         string                  `json:"overview"`
	ImdbID           string                  `json:"imdb_id"`
	OriginalLanguage string                  `json:"original_language"`
}

type TmdbCastResponse struct {
	Cast []struct {
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
		Character   string `json:"character"`
	} `json:"cast"`
	Crew []struct {
		Name        string `json:"name"`
		ProfilePath string `json:"profile_path"`
		Job         string `json:"job"`
	} `json:"crew"`
}

func MapTmdbToFilm(tmdbFilmData, tmdbCastData []byte, screeningDate time.Time) (*db.Film, []db.Director, []db.CastMember, error) {
	var tmdbFilm TmdbFilm
	var tmdbCast TmdbCastResponse

	if err := json.Unmarshal(tmdbFilmData, &tmdbFilm); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse film data: %w", err)
	}
	if err := json.Unmarshal(tmdbCastData, &tmdbCast); err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse cast data: %w", err)
	}

	// Map cast (first 4)
	var cast []db.CastMember
	for i, c := range tmdbCast.Cast {
		if i >= 4 {
			break
		}
		character := c.Character
		if strings.TrimSpace(character) == "" {
			character = "Unknown Character"
		}
		cast = append(cast, db.CastMember{
			Name:      c.Name,
			ImgUrl:    "https://image.tmdb.org/t/p/w500" + c.ProfilePath,
			Character: character,
		})
	}

	// Map directors (crew with job == "Director")
	var directors []db.Director
	for _, crew := range tmdbCast.Crew {
		if crew.Job == "Director" {
			directors = append(directors, db.Director{
				Name:   crew.Name,
				ImgUrl: "https://image.tmdb.org/t/p/w500" + crew.ProfilePath,
			})
		}
	}

	// Map genres
	var genres []db.Genre
	for _, g := range tmdbFilm.Genres {
		genres = append(genres, db.Genre{Name: g.Name})
	}

	// Map language
	language := ""
	if len(tmdbFilm.SpokenLanguages) > 0 {
		language = tmdbFilm.SpokenLanguages[0].Name
	}

	// Parse release date
	releaseDate, _ := time.Parse("2006-01-02", tmdbFilm.ReleaseDate)

	// Map film
	film := &db.Film{
		Name:          tmdbFilm.Title,
		ImgURL:        "https://image.tmdb.org/t/p/w500" + tmdbFilm.PosterPath,
		ReleaseDate:   releaseDate,
		ScreeningDate: screeningDate,
		Duration:      tmdbFilm.Runtime,
		Language:      language,
		Overview:      tmdbFilm.Overview,
		Genres:        genres,
		IMDbURL:       fmt.Sprintf("https://www.imdb.com/title/%s/", tmdbFilm.ImdbID),
		Cast:          cast,
		Directors:     directors,
		Reviews:       []db.Review{},
	}

	return film, directors, cast, nil
}
