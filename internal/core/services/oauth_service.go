package services

import "qropen-backend/internal/core/domain"

type OAuthService struct{}

func NewOAuthService() *OAuthService {
	return &OAuthService{}
}

func (s *OAuthService) GetGoogleAuthURL() (string, error) {
	return "", nil
}

func (s *OAuthService) HandleGoogleCallback(code string) (domain.User, error) {
	return domain.User{}, nil
}
