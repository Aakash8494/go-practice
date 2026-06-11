package service

import (
	"errors"
	"myapp/domain"
)

type UserService struct {
	repo domain.UserRepository // It only trusts the interface!
}

// Constructor to inject the database
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Business Logic: Creating a user
func (s *UserService) RegisterUser(user *domain.User) error {
	// Simple validation
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are strictly required")
	}

	// It calls the interface method. It doesn't care if it's Mongo or SQL!
	return s.repo.CreateUser(user)
}

// Business Logic: Getting a user
func (s *UserService) FetchUser(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.GetUser(id)
}
