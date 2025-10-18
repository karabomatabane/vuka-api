package db

type Source struct {
	Model
	Name       string `json:"name" gorm:"uniqueIndex:unique_source_name_website"`
	WebsiteUrl string `json:"websiteUrl" gorm:"uniqueIndex:unique_source_name_website"`
	RssFeedUrl string `json:"rssFeedUrl"`
}
