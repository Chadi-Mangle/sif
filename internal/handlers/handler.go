package handlers

import (
	"context"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/token"
)

type Handler struct {
	ctx     context.Context
	queries models.Queries
	token   *token.JWTHandler
}

func NewHandler(ctx context.Context, queries models.Queries, secretKey string) *Handler {
	return &Handler{
		ctx:     ctx,
		queries: queries,
		token:   token.NewJTWHandler(secretKey),
	}
}
