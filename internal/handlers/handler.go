package handlers

import (
	"context"
	"net/http"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/auth"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/internal/token"
	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
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

// GetDashboard affiche le dashboard avec les informations de l'utilisateur connecté
func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	// Récupérer les infos utilisateur depuis le cookie JWT
	accessToken, err := auth.GetAccessToken(r)
	if err != nil {
		// Pas connecté → redirection login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Vérifier et décoder le token
	claims, err := h.token.VerifyToken(accessToken)
	if err != nil {
		// Token invalide/expiré → redirection login
		// TODO: Implémenter la logique de refresh token ici
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Rendu du template avec les données utilisateur
	component := templates.Dashboard(claims.FirstName, claims.LastName, claims.IsAdmin)
	component.Render(r.Context(), w)
}
