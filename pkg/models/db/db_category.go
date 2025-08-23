package db

type Category struct {
	Model
	Name     string     `json:"name"`
	Articles []*Article `gorm:"many2many:article_categories;"`
}
