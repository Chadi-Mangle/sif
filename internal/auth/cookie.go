package auth

import (
	"net/http"
)

const (
	AccessTokenCookie  = "access_token"
	RefreshTokenCookie = "refresh_token"
)

func SetAuthCookies(w http.ResponseWriter, accessToken, refreshToken string) {
	// Cookie pour l'access token (httpOnly, sécurisé, courte durée)
	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenCookie,
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true en production avec HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   15 * 60, // 15 minutes
	})

	// Cookie pour le refresh token (httpOnly, sécurisé, longue durée)
	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookie,
		Value:    refreshToken,
		Path:     "/auth/refresh", // Limité au endpoint de refresh
		HttpOnly: true,
		Secure:   false, // true en production avec HTTPS
		SameSite: http.SameSiteStrictMode,
		MaxAge:   7 * 24 * 60 * 60, // 7 jours
	})
}

func GetAccessToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(AccessTokenCookie)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func GetRefreshToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(RefreshTokenCookie)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ClearAuthCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     AccessTokenCookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     RefreshTokenCookie,
		Value:    "",
		Path:     "/auth/refresh",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
