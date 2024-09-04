package ports

import "qropen-backend/internal/core/domain"

type OAuthService interface {
	GetGoogleAuthURL() (string, error)
	HandleGoogleCallback(code string) (domain.User, error)
}
