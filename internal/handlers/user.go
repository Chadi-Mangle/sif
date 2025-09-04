package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/models"
	"github.com/Chadi-Mangle/templ-hmr-setup/package/utils"
	"github.com/Chadi-Mangle/templ-hmr-setup/templates"
)

func (h *Handler) GetSignUp(w http.ResponseWriter, r *http.Request) {
	composant := templates.Register()
	composant.Render(r.Context(), w)
}

func (h *Handler) PostSingUp(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	password := r.FormValue("password")

	hashedPassword, errPassword := utils.HashPassword(password)
	if errPassword != nil {
		return //Erreur lors de l'encodage du mot de passe
	}

	createUser := models.CreateUserParams{
		FirstName:    firstName,
		LastName:     lastName,
		HashPassword: hashedPassword,
		IsActivated:  false,
		HasPaid:      false,
		IsAdmin:      false,
	}

	// Faire une logique d'envoie de mail sur l'adresse académique pour
	// que les gens hors de l'ecole ne puisse pas s'inscirire

	_, errUser := h.queries.CreateUser(h.ctx, createUser)
	if errUser != nil {
		return //Erreur lors de la création de l'utilisateur
	}

}

func (h *Handler) GetSignIn(w http.ResponseWriter, r *http.Request) {
	composant := templates.Login()
	composant.Render(r.Context(), w)
}

func (h *Handler) PostSignIn(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	user := models.GetUserPasswordByNameParams{
		FirstName: firstName,
		LastName:  lastName,
	}

	hashedPassword, errUser := h.queries.GetUserPasswordByName(h.ctx, user)
	if errUser != nil {
		// Voir de faire un goto pour handler les erreurs en html
		return // User not fond en gros
	}

	password := r.FormValue("password")

	errPassword := utils.CheckPassword(password, hashedPassword)
	if errPassword != nil {
		return // Mauvais mot de passe
	}

	fmt.Printf("%s %s connecté", firstName, lastName)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.CreateUserParams
	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "error decoding request body", http.StatusBadRequest)
		return
	}

}
