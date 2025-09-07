package token

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTHandler struct {
	secretKey string
}

type RefreshToken struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
	Revoked   bool
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	AccessClaims *UserClaims
}

func NewJTWHandler(secretKey string) *JWTHandler {
	return &JWTHandler{secretKey}
}

func (h *JWTHandler) CreateToken(firstName string, lastName string, isAdmin bool, duration time.Duration) (string, *UserClaims, error) {
	claims, err := NewUserClaims(firstName, lastName, isAdmin, duration)
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

func (h *JWTHandler) CreateTokenPair(firstName string, lastName string, isAdmin bool) (*TokenPair, error) {
	// Créer l'access token (courte durée - 15 minutes)
	accessToken, claims, err := h.CreateToken(firstName, lastName, isAdmin, 15*time.Minute)
	if err != nil {
		return nil, err
	}

	// Créer le refresh token (longue durée - 7 jours)
	refreshToken, err := h.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessClaims: claims,
	}, nil
}

func (h *JWTHandler) generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (h *JWTHandler) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		_, isTrue := t.Method.(*jwt.SigningMethodHMAC)
		if !isTrue {
			return nil, errors.New("mauvaise méthode d'encryption")
		}

		return []byte(h.secretKey), nil
	})

	if err != nil {
		return nil, errors.New("erreur dans le token")
	}

	claims, isTrue := token.Claims.(*UserClaims)
	if !isTrue {
		return nil, errors.New("erreur avec les claims token")
	}

	return claims, nil
}
