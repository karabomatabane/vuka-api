package db

type Source struct {
	Model
	Name       string `json:"name"`
	WebsiteUrl string `json:"websiteUrl"`
	RssFeedUrl string `json:"rssFeedUrl"`
}
