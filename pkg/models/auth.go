package models

type RegisterBody struct {
	Username string `json:"username" validate:"required, min=3"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   string `json:"roleId" validate:"required"`
}

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Username    string `json:"username"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
}
