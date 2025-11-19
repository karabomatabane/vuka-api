package utils

import (
	"math"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// PaginationParams represents pagination parameters
type PaginationParams struct {
	Page     int
	PageSize int
}

// PaginationResult represents paginated response metadata
type PaginationResult struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Data       interface{}      `json:"data"`
	Pagination PaginationResult `json:"pagination"`
}

// GetPaginationParams extracts and validates pagination parameters from query strings
func GetPaginationParams(pageStr, pageSizeStr string) PaginationParams {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10 // default page size
	} else if pageSize > 100 {
		pageSize = 100 // max page size
	}

	return PaginationParams{
		Page:     page,
		PageSize: pageSize,
	}
}

// CalculateOffset calculates the database offset from page and pageSize
func (p PaginationParams) CalculateOffset() int {
	return (p.Page - 1) * p.PageSize
}

// CreatePaginationResult creates pagination metadata
func CreatePaginationResult(page, pageSize int, totalItems int64) PaginationResult {
	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	return PaginationResult{
		Page:       page,
		PageSize:   pageSize,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
}

func GenerateTokenString(userId uuid.UUID, roleId uuid.UUID, roleName string, expDate time.Time) (string, error) {
	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"roleId": roleId,
		"role":   roleName,
		"exp":    expDate.Unix(), // expires in 24 hours
	})

	// Sign the token and get the complete encoded token as a string
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}
