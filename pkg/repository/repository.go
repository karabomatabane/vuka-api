package repository

import (
	"vuka-api/pkg/repository/contracts"
	"vuka-api/pkg/repository/implementations"

	"gorm.io/gorm"
)

type Repositories struct {
	Article  contracts.ArticleRepository
	User     contracts.UserRepository
	Role     contracts.RoleRepository
	Source   contracts.SourceRepository
	Category contracts.CategoryRepository
	Directory contracts.DirectoryRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Article:  implementations.NewArticleRepository(db),
		User:     implementations.NewUserRepository(db),
		Role:     implementations.NewRoleRepository(db),
		Source:   implementations.NewSourceRepository(db),
		Category: implementations.NewCategoryRepository(db),
		Directory: implementations.NewDirectoryRepository(db),
	}
}
