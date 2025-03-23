package services

import (
	"errors"
	"freshers-bootcamp-last/models"
	"freshers-bootcamp-last/repositories"
	"freshers-bootcamp-last/utils"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user *models.User) error {
	// Check if the username is already taken
	existingUser, err := s.repo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}

	// Hash the password before saving
	user.Password = utils.HashPassword(user.Password)

	// Set default role to "user" if not provided
	if user.Role == "" {
		user.Role = "user"
	}

	// Save the user
	return s.repo.Create(user)
}

func (s *AuthService) RegisterAdmin(user *models.User) error {
	// Check if the username is already taken
	existingUser, err := s.repo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return errors.New("username already taken")
	}

	// Hash the password before saving
	user.Password = utils.HashPassword(user.Password)

	// Set role to "admin"
	user.Role = "admin"

	// Save the user
	return s.repo.Create(user)
}

func (s *AuthService) Login(username, password string) (string, error) {
	// Find the user by username
	user, err := s.repo.FindByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("invalid username or password")
	}

	// Verify the password
	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("invalid username or password")
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
