package session

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	sessionKey = "session"
)

var (
	clientEntryPoint string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

	if clientEntryPoint = os.Getenv("CLIENT_ENTRY_POINT"); clientEntryPoint == "" {
		log.Fatalln("Environment value CLIENT_ENTRY_POINT is empty.")
	}
}
