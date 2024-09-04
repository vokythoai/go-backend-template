package jwt

import (
	"qropen-backend/internal/core/ports"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtAdapter struct {
	secretKey []byte
}

func NewJWTAdapter(secretKey string) ports.JWTAdapter {
	return &jwtAdapter{secretKey: []byte(secretKey)}
}

func (a *jwtAdapter) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(a.secretKey)
}

func (a *jwtAdapter) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}

	return "", jwt.ErrSignatureInvalid
}
