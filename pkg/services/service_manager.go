package services

import (
	"vuka-api/pkg/repository"

	"gorm.io/gorm"
)

type Services struct {
	Film *FilmService
	User *UserService
	Auth *AuthService
	// Review *ReviewService
}

func NewServices(db *gorm.DB) *Services {
	repos := repository.NewRepositories(db)

	return &Services{
		Film: NewFilmService(repos),
		User: NewUserService(repos),
		Auth: NewAuthService(repos),
		// Initialize other services here
	}
}
