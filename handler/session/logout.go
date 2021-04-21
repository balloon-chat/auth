package session

import (
	"github.com/baloon/go/auth/env"
	"net/http"
)

// Logout セッション情報を削除する
func Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", clientEntryPoint)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	switch r.Method {
	case http.MethodPost:
		http.SetCookie(w, &http.Cookie{
			Name:     sessionKey,
			Value:    "",
			MaxAge:   -1, // Cookieを削除する
			HttpOnly: true,
			Secure:   !env.DEBUG,
			Path:     "/",
		})
	}

	w.WriteHeader(http.StatusOK)
}
