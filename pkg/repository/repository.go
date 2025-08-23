package repository

import (
	"vuka-api/pkg/repository/contracts"
	"vuka-api/pkg/repository/implementations"

	"gorm.io/gorm"
)

type Repositories struct {
	Film       contracts.FilmRepository
	Genre      contracts.GenreRepository
	Director   contracts.DirectorRepository
	CastMember contracts.CastMemberRepository
	Review     contracts.ReviewRepository
	User       contracts.UserRepository
	Auth       contracts.AuthRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Film:       implementations.NewFilmRepository(db),
		Genre:      implementations.NewGenreRepository(db),
		Director:   implementations.NewDirectorRepository(db),
		CastMember: implementations.NewCastMemberRepository(db),
		Review:     implementations.NewReviewRepository(db),
		User:       implementations.NewUserRepository(db),
		Auth:       implementations.NewAuthRepository(db),
	}
}
