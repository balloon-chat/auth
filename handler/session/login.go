package session

import (
	"context"
	"encoding/json"
	"github.com/baloon/go/auth/app/infrastructure/firebase"
	"github.com/baloon/go/auth/env"
	"log"
	"net/http"
	"time"
)

type Request struct {
	IdToken string `json:"idToken"`
}

// Login Firestoreが発行するトークンを用いてセッションを作成する
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", clientEntryPoint)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	switch r.Method {
	case http.MethodOptions:
		return
	case http.MethodPost:
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("error while decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 有効期限: 5日
	expiresIn := 24 * time.Hour * 5

	client, err := firebase.NewFirebaseAuthClient(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	decoded, err := client.VerifyIDToken(r.Context(), request.IdToken)
	if err != nil {
		log.Println("Invalid ID token", err)
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}

	// 最終ログインが5分以内でなければ、再ログインを要求
	if time.Now().Unix()-decoded.AuthTime > 5*60 {
		http.Error(w, "Recent sign-in required", http.StatusUnauthorized)
		return
	}

	// セッションCookieを作成
	cookie, err := client.SessionCookie(r.Context(), request.IdToken, expiresIn)
	if err != nil {
		http.Error(w, "Failed to create session cookie", http.StatusInternalServerError)
		log.Println("Failed to create session cookie:", err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     sessionKey,
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   !env.DEBUG,
		Path:     "/",
	})

	w.WriteHeader(http.StatusOK)
}
