package ports

type AuthService interface {
	Login(username, password string) (string, error)
	ValidateToken(token string) (string, error)
}
