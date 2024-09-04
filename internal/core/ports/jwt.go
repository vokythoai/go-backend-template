package ports

type JWTAdapter interface {
	GenerateToken(username string) (string, error)
	ValidateToken(token string) (string, error)
}
