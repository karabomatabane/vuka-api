package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

func authenticateToken(r *http.Request) (*jwt.Token, jwt.Claims, error) {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	tokenString := strings.Split(authHeader, " ")[1]
	// Check if the tokenString is valid
	if tokenString == "" {
		return nil, nil, jwt.ErrSignatureInvalid
	}

	// Decode the token using the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, nil, err
	}

	// Attach user information to the context
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, err
	}
	
	return token, claims, nil
}

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := authenticateToken(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}
		
		// Add claims to the request context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func VerifyTokenAndAdminFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := authenticateToken(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		// Check if the user has admin role
		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok || role != "admin" {
			http.Error(w, "Admin resource! Access denied.", http.StatusForbidden)
			return
		}

		// Add claims to the request context
		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func VerifyTokenAndAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := authenticateToken(r)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		role, ok := claims.(jwt.MapClaims)["role"].(string)
		if !ok || role != "admin" {
			http.Error(w, "Admin resource! Access denied.", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

