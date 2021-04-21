package main

import (
	"github.com/baloon/go/auth/handler/oauth/google"
	"github.com/baloon/go/auth/handler/session"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/oauth/google", google.Oauth)
	http.HandleFunc("/oauth/google/callback", google.OauthCallback)
	http.HandleFunc("/oauth/google/result", google.GetOauthResult)

	http.HandleFunc("/session/login", session.Login)
	http.HandleFunc("/session/logout", session.Logout)
	http.HandleFunc("/profile", session.GetProfile)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
