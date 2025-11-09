package services

import (
	"vuka-api/pkg/models/db"
	"vuka-api/pkg/models/user"
	"vuka-api/pkg/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repos *repository.Repositories
}

func NewUserService(repos *repository.Repositories) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) GetUserByID(id string) (*db.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.repos.User.GetByID(userId)
}

func (s *UserService) GetUserByUsername(username string) (*db.User, error) {
	return s.repos.User.GetByUsername(username)
}

func (s *UserService) GetAllUsers() ([]db.User, error) {
	return s.repos.User.GetAll()
}

func (s *UserService) UpdateUserRole(body user.UpdateUserRoleBody) (*db.User, error) {
	u, err := s.GetUserByID(body.UserID)
	if err != nil {
		return nil, err
	}
	u.RoleID = body.RoleID
	if err := s.repos.User.Update(u); err != nil {
		return nil, err
	}
	updatedUser, err := s.GetUserByID(body.UserID)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (s *UserService) CreateUser(user *db.User) error {
	return s.repos.User.Create(user)
}

func (s *UserService) UpdateUser(user *db.User) error {
	return s.repos.User.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repos.User.Delete(userID)
}
