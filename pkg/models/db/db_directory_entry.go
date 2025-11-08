package db

import "github.com/google/uuid"

type DirectoryEntry struct {
	Model
	Name        string            `json:"name"`
	Description string            `json:"description"`
	ContactInfo []ContactInfo     `json:"contactInfo" gorm:"foreignKey:DirectoryEntryID"`
	WebsiteUrl  string            `json:"websiteUrl"`
	EntryType   string            `json:"entryType"`
	CategoryID  uuid.UUID         `json:"categoryId" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category    DirectoryCategory `json:"category"`
}
