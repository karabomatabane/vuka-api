package db

import (
	"github.com/google/uuid"
	"time"
)

type Article struct {
	Model
	Title       string      `json:"title,omitempty"`
	Language    string      `json:"language,omitempty"`
	OriginalUrl string      `json:"originalUrl,omitempty"`
	ContentBody string      `json:"contentBody,omitempty"`
	PublishedAt time.Time   `json:"publishedAt"`
	IsFeatured  bool        `json:"isFeatured,omitempty"`
	SourceID    uuid.UUID   `json:"sourceId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Source      Source      `json:"source"`
	RegionID    string      `json:"regionID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Region      Region      `json:"region"`
	Categories  []*Category `gorm:"many2many:article_categories;"`
}
