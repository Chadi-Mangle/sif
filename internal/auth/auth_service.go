package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/token"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService struct {
	queries    *models.Queries
	jwtHandler *token.JWTHandler
}

func NewAuthService(queries *models.Queries, jwtHandler *token.JWTHandler) *AuthService {
	return &AuthService{
		queries:    queries,
		jwtHandler: jwtHandler,
	}
}

func (s *AuthService) Login(ctx context.Context, firstName, lastName string) (*token.TokenPair, error) {
	user, err := s.queries.GetUserByName(ctx, models.GetUserByNameParams{
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return nil, err
	}

	tokenPair, err := s.jwtHandler.CreateTokenPair(firstName, lastName, user.IsAdmin)
	if err != nil {
		return nil, err
	}

	refreshTokenID := uuid.New()
	expiresAt := pgtype.Timestamp{
		Time:  time.Now().Add(7 * 24 * time.Hour),
		Valid: true,
	}
	_, err = s.queries.CreateRefreshToken(ctx, models.CreateRefreshTokenParams{
		ID:        refreshTokenID.String(),
		FirstName: firstName,
		LastName:  lastName,
		Token:     tokenPair.RefreshToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshTokenString string) (*token.TokenPair, error) {
	refreshToken, err := s.queries.GetRefreshTokenByToken(ctx, refreshTokenString)
	if err != nil {
		return nil, err
	}

	if refreshToken.ExpiresAt.Valid && time.Now().After(refreshToken.ExpiresAt.Time) {
		return nil, errors.New("refresh token expiré")
	}

	user, err := s.queries.GetUserByName(ctx, models.GetUserByNameParams{
		FirstName: refreshToken.FirstName,
		LastName:  refreshToken.LastName,
	})
	if err != nil {
		return nil, err
	}

	err = s.queries.RevokeRefreshToken(ctx, refreshToken.ID)
	if err != nil {
		return nil, err
	}

	return s.Login(ctx, user.FirstName, user.LastName)
}

func (s *AuthService) Logout(ctx context.Context, refreshTokenString string) error {
	refreshToken, err := s.queries.GetRefreshTokenByToken(ctx, refreshTokenString)
	if err != nil {
		return err
	}

	return s.queries.RevokeRefreshToken(ctx, refreshToken.ID)
}
