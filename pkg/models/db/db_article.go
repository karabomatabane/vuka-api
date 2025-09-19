package db

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	Model
	Title       string      `json:"title,"`
	Language    string      `json:"language,"`
	OriginalUrl string      `json:"originalUrl" gorm:"index"`
	ContentBody string      `json:"contentBody"`
	PublishedAt time.Time   `json:"publishedAt"`
	IsFeatured  bool        `json:"isFeatured"`
	SourceID    *uuid.UUID  `json:"sourceId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Source      Source      `json:"source"`
	RegionID    *string     `json:"regionID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Region      Region      `json:"region"`
	Categories  []*Category `gorm:"many2many:article_categories;" json:"categories"`
}
