package models

import "github.com/google/uuid"

type DirectoryOverviewResponse struct {
	Categories   []DirectoryCategoryResponse `json:"categories"`
	Personalized PersonalisedData            `json:"personalised"`
}

type PersonalisedData struct {
	Pinned []DirectoryCategoryResponse `json:"pinned"`
	Recent []DirectoryCategoryResponse `json:"recent"`
}

type DirectoryCategoryResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	TotalEntries int64     `json:"totalEntries"`
}
