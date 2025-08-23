package models

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/models/user"
)

type RegisterBody struct {
	db.User
	AccountCode string `json:"accountCode"`
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Username    string    `json:"username"`
	Role        user.Role `json:"role"`
	AccessToken string    `json:"accessToken"`
}

type AccountCode struct {
	Code      string `json:"code"`
	DaysValid int32  `json:"daysValid"`
}
