package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
	"vuka-api/pkg/models/user"
)

func GenerateTokenString(userId uuid.UUID, userRole user.Role, expDate time.Time) (string, error) {
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userId,
		"role": userRole,
		"exp":  expDate.Unix(), // expires in 24 hours
	})

	// Sign the token and get the complete encoded token as a string
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}
