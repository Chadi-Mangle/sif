package token

import (
	"errors"
	"time"

	"github.com/Chadi-Mangle/templ-hmr-setup/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	jwt.RegisteredClaims
}

func NewUserClaims(firstName string, lastName string, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("Erreur lors de la création du token ID")
	}

	return &UserClaims{
		FirstName: firstName,
		LastName:  lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			Subject:   utils.GetEmailAddress(firstName, lastName),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}, nil

}
