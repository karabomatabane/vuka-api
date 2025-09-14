package services

import (
	"vuka-api/pkg/repository"

	"gorm.io/gorm"
)

type Services struct {
	Article *ArticleService
	User    *UserService
	Auth    *AuthService
	Role    *RoleService
	Rss     *RssService
	Source  *SourceService
}

func NewServices(db *gorm.DB) *Services {
	repos := repository.NewRepositories(db)

	return &Services{
		Article: NewArticleService(repos),
		User:    NewUserService(repos),
		Auth:    NewAuthService(repos),
		Role:    NewRoleService(repos),
		Rss:     NewRssService(),
		Source:  NewSourceService(repos),
	}
}
