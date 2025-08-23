package services

import (
	"errors"
	"fmt"
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
	ac, err := s.repos.Auth.GetAccountCode()
	if err != nil {
		return nil, err
	}
	if ac.Code != body.AccountCode {
		return nil, errors.New("invalid account code")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dbUser := &db.User{
		Username: body.Username,
		Password: string(password),
		Role:     body.Role,
	}

	err = s.repos.User.Create(dbUser)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateTokenString(dbUser.ID, dbUser.Role, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}
	user := models.AuthResponse{
		Username:    dbUser.Username,
		Role:        dbUser.Role,
		AccessToken: token,
	}
	return &user, nil
}

func (s *AuthService) Login(body models.LoginBody) (*models.AuthResponse, error) {
	dbUser, err := s.repos.User.GetByUsername(body.Username)
	if err != nil {
		return nil, err
	}
	if dbUser == nil {
		return nil, errors.New(httpx.InvalidCredentials)
	}

	// Compare the incoming password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(body.Password)); err != nil {
		fmt.Printf("Password comparison failed for user %s: %v\n", body.Username, err)
		return nil, errors.New(httpx.InvalidCredentials)
	}

	token, err := utils.GenerateTokenString(dbUser.ID, dbUser.Role, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}
	user := models.AuthResponse{
		Username:    dbUser.Username,
		Role:        dbUser.Role,
		AccessToken: token,
	}

	return &user, nil
}

func (s *AuthService) CreateAccountCode(body models.AccountCode) (*db.AccountCode, error) {
	activeCode, err := s.repos.Auth.GetAccountCode()
	if err != nil {
		return nil, err
	}
	if activeCode != nil && activeCode.Code != "" {
		if err = s.repos.Auth.DeleteAccountCode(activeCode); err != nil {
			return nil, err
		}
	}

	// there's no active code ∴ create one.
	activeCode.Code = body.Code
	hoursValid := 24 * time.Duration(body.DaysValid)
	activeCode.ExpirationDate = time.Now().Add(time.Hour * hoursValid)
	err = s.repos.Auth.CreateAccountCode(activeCode)
	if err != nil {
		return nil, err
	}

	return activeCode, nil
}

func (s *AuthService) GetAccountCode() (*db.AccountCode, error) {
	activeCode, err := s.repos.Auth.GetAccountCode()
	if err != nil {
		return nil, err
	}

	return activeCode, nil
}

func (s *AuthService) DeleteAccountCode() error {
	activeCode, err := s.repos.Auth.GetAccountCode()
	if err != nil {
		return err
	}
	return s.repos.Auth.DeleteAccountCode(activeCode)
}
