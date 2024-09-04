package services

import (
	"errors"
	"qropen-backend/internal/core/domain"
	"qropen-backend/internal/core/ports"
)

type authService struct {
	repository ports.Repositories
	jwtAdapter ports.JWTAdapter
}

func NewAuthService(repository ports.Repositories, jwtAdapter ports.JWTAdapter) ports.AuthService {
	return &authService{repository: repository, jwtAdapter: jwtAdapter}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.repository.UserRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", errors.New("invalid credentials")
	}
	return s.jwtAdapter.GenerateToken(username)
}

func (s *authService) ValidateToken(token string) (string, error) {
	return s.jwtAdapter.ValidateToken(token)
}

func (s *authService) CreateUser(user domain.User) error {
	return s.repository.UserRepo.CreateUser(user)
}
