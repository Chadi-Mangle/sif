package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
)

func (h *Handler) SignUp() {

}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	composant := templates.Login()
	composant.Render(r.Context(), w)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

	if _, err := h.queries.CreateUser(h.ctx, createUser); err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}
}
