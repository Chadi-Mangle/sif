package handlers

import (
	"context"
	"encoding/json"
	"net/http"

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

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if user, err := h.queries.CreateUser(h.ctx, createUser); err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	res := user
}
