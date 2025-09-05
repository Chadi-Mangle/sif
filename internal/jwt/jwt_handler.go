package token

import "time"

type JWTHandler struct {
	secretKey string
}

func NewJTWHandler(secretKey string) *JWTHandler {
	return &JWTHandler{secretKey}
}

func (h *JWTHandler) CreateToken(firstName string, lastName string, duration time.Duration) (string, *UserClaims, error) {

}
