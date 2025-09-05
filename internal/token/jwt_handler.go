package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTHandler struct {
	secretKey string
}

func NewJTWHandler(secretKey string) *JWTHandler {
	return &JWTHandler{secretKey}
}

func (h *JWTHandler) CreateToken(firstName string, lastName string, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(firstName, lastName, duration)
	if err != nil {
		return "", nil, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.secretKey))
	if err != nil {
		return "", nil, err
	}

	return tokenString, claims, nil
}

func (h *JWTHandler) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		_, isTrue := t.Method.(*jwt.SigningMethodHMAC)
		if !isTrue {
			return nil, errors.New("Mauvaise méthode d'encryption")
		}

		return []byte(h.secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Erreur dans le token")
	}

	claims, isTrue := token.Claims.(*UserClaims)
	if !isTrue {
		return nil, errors.New("Erreur dans les calims token")
	}

	return claims, nil
}
