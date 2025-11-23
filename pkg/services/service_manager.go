package services

import (
	"vuka-api/pkg/repository"

	"gorm.io/gorm"
)

type Services struct {
	Article    *ArticleService
	User       *UserService
	Auth       *AuthService
	Role       *RoleService
	Rss        *RssService
	Source     *SourceService
	Directory  *DirectoryService
	Cron       *CronService
	Category   *CategoryService
	Permission *PermissionService
	Newsletter *NewsletterService
}

func NewServices(db *gorm.DB) *Services {
	repos := repository.NewRepositories(db)

	articleService := NewArticleService(repos)
	sourceService := NewSourceService(repos)
	categoryService := NewCategoryService(repos.Category)
	rssService := NewRssService(articleService, categoryService)
	directoryService := NewDirectoryService(repos.Directory)
	newsletterService := NewNewsletterService(repos)

	return &Services{
		Article:    articleService,
		User:       NewUserService(repos),
		Auth:       NewAuthService(repos),
		Role:       NewRoleService(repos),
		Rss:        rssService,
		Source:     sourceService,
		Cron:       NewCronService(rssService, sourceService, newsletterService),
		Category:   categoryService,
		Directory:  directoryService,
		Permission: NewPermissionService(repos),
		Newsletter: newsletterService,
	}
}
