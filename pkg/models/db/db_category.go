package db

type Category struct {
	Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Articles    []*Article `json:"articles" gorm:"many2many:article_categories;"`
}
