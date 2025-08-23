package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

func GenerateTokenString(userId, roleId uuid.UUID, expDate time.Time) (string, error) {
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"roleId": roleId,
		"exp":    expDate.Unix(), // expires in 24 hours
	})

	// Sign the token and get the complete encoded token as a string
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}
