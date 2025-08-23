package repository

import (
	"vuka-api/pkg/repository/contracts"
	"vuka-api/pkg/repository/implementations"

	"gorm.io/gorm"
)

type Repositories struct {
	Article contracts.ArticleRepository
	User    contracts.UserRepository
	Role    contracts.RoleRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Article: implementations.NewArticleRepository(db),
		User:    implementations.NewUserRepository(db),
		Role:    implementations.NewRoleRepository(db),
	}
}
