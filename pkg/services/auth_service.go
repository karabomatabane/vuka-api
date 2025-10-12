package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
	"vuka-api/pkg/httpx"
	"vuka-api/pkg/models"
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/repository"
	"vuka-api/pkg/utils"
)

type AuthService struct {
	repos *repository.Repositories
}

func NewAuthService(repos *repository.Repositories) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) Register(body models.RegisterBody) (*models.AuthResponse, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	roleId, err := uuid.Parse(body.RoleID)
	if err != nil {
		return nil, err
	}

	dbUser := &db.User{
		Username:     body.Username,
		PasswordHash: string(password),
		RoleID:       roleId,
	}

	err = s.repos.User.Create(dbUser)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateTokenString(dbUser.ID, dbUser.RoleID, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}
	dbUserWithRole, err := s.repos.User.GetByID(dbUser.ID)
	if err != nil {
		return nil, err
	}
	user := models.AuthResponse{
		Username:    dbUserWithRole.Username,
		Role:        dbUserWithRole.Role.Name,
		AccessToken: token,
	}
	return &user, nil
}

func (s *AuthService) Login(body models.LoginBody) (*models.AuthResponse, error) {
	dbUser, err := s.repos.User.GetByUsername(body.Username)
	if err != nil {
		// TODO: Return err for reporting, then controller will send errors.New(httpx.InvalidCredentials) to client
		return nil, errors.New(httpx.InvalidCredentials)
	}
	if dbUser == nil {
		return nil, errors.New(httpx.InvalidCredentials)
	}

	// Compare the incoming password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(body.Password)); err != nil {
		fmt.Printf("Password comparison failed for user %s: %v\n", body.Username, err)
		return nil, errors.New(httpx.InvalidCredentials)
	}

	token, err := utils.GenerateTokenString(dbUser.ID, dbUser.RoleID, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, errors.New(httpx.InvalidCredentials)
	}
	user := models.AuthResponse{
		Username:    dbUser.Username,
		Role:        dbUser.Role.Name,
		AccessToken: token,
	}

	return &user, nil
}
