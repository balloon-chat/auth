package session

import (
	"context"
	"encoding/json"
	"github.com/baloon/go/auth/app/infrastructure/firebase"
	"log"
	"net/http"
)

type Response struct {
	// Firebaseログインで用いられるユーザー識別子
	LoginId string `json:"loginId"`
}

// GetProfile Firebaseによって作成されたセッション情報を用いて、ユーザーのプロフィールを取得
func GetProfile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", clientEntryPoint)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	switch r.Method {
	case http.MethodOptions:
		w.Header()
		return
	case http.MethodGet:
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cookie, err := r.Cookie(sessionKey)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("error while getting cookie:", err)
		return
	}

	client, err := firebase.NewFirebaseAuthClient(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	decoded, err := client.VerifySessionCookie(r.Context(), cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("error while verifying session cookie", err)
		return
	}

	res := Response{
		LoginId: decoded.UID,
	}
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error while encoding response", err)
		return
	}
}
