package db

import (
	"github.com/google/uuid"
	"time"
)

type UserDirectoryMeta struct {
	Model
	UserID       uuid.UUID      `json:"userId"`
	User         User           `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DirectoryID  uuid.UUID      `json:"directoryId"`
	Directory    DirectoryCategory `json:"directory" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Pinned       bool           `json:"pinned"`
	LastAccessed time.Time      `json:"lastAccessed"`
}
