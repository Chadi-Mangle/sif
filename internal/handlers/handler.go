package handlers

import (
	"context"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
)

type Handler struct {
	ctx     context.Context
	queries models.Queries
}

func NewHandler(ctx context.Context, queries models.Queries) *Handler {
	return &Handler{
		ctx:     ctx,
		queries: queries,
	}
}
