package httpx

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
	"vuka-api/pkg/models/user"
)

// ParseBody decodes JSON into the given struct
func ParseBody(r *http.Request, x any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // optional strict mode
	return dec.Decode(x)
}

// WriteJSON writes a JSON response with the given status code
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		WriteErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
}

// WriteErrorJSON writes an error message as JSON with the given status code
func WriteErrorJSON(w http.ResponseWriter, message string, status int) {
	WriteJSON(w, status, map[string]string{"error": message})
}

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

func GetFromTmdb(url string) (*http.Response, error) {
	token := os.Getenv("TMDB_ACCESS_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TMDB_ACCESS_TOKEN not set")
	}
	req, err := http.NewRequest(HTTP_GET, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return http.DefaultClient.Do(req)
}
