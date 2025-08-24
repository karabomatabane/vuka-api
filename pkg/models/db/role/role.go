package role

import "github.com/google/uuid"

type Response struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Permissions map[string][]string `json:"permissions"`
}
