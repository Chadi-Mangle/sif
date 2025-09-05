package handlers

import (
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
		fmt.Printf("password %v", errPassword)
		return // Erreur lors de l'encodage du mot de passe
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
	// que les gens hors de l'ecole ne puisse pas s'inscrire

	_, errUser := h.queries.CreateUser(h.ctx, createUser)
	if errUser != nil {
		fmt.Printf("%v", errUser)
		return // Erreur lors de la création de l'utilisateur
	}

	fmt.Printf("Utilisateur créé\n")
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

	hashedPassword, err := h.queries.GetUserPasswordByName(h.ctx, user)
	if err != nil {
		fmt.Printf("Erreur pour trouver le mdp : %v", err)
		// Voir de faire un goto pour handler les erreurs en html
		return // User not fond en gros
	}

	password := r.FormValue("password")

	err = utils.CheckPassword(password, hashedPassword)
	if err != nil {
		fmt.Printf("%v", err)
		return // Mauvais mot de passe
	}

	fmt.Printf("%s %s est connecté", firstName, lastName)

	// accessToken, _, err := h.token.CreateToken(firstName, lastName, 15*time.Minute)
	// if err != nil {
	// 	return // Erreur lors de la création du token
	// }

	// http.SetCookie()
}
