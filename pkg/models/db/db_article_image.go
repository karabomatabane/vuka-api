package db

import (
	"github.com/google/uuid"
)

type ArticleImage struct {
	Model
	ArticleID uuid.UUID `gorm:"type:uuid;index"`
	IsMain    bool      `json:"isMain"`
	URL       string    `json:"url"`
	AltText   string    `json:"altText,omitempty"`
}
