package main

import (
	"github.com/baloon/go/oauth/handler/auth"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/oauth/google", auth.OauthGoogle)
	http.HandleFunc("/oauth/google/callback", auth.OauthCallbackGoogle)
	http.HandleFunc("/oauth/google/token", auth.GetAccessToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}
