package main

import (
	"github.com/baloon/go/auth/handler/oauth/google"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/oauth/google", google.Oauth)
	http.HandleFunc("/oauth/google/callback", google.OauthCallback)
	http.HandleFunc("/oauth/google/token", google.GetAccessToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
