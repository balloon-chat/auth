package session

import (
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
	if clientEntryPoint = os.Getenv("CLIENT_ENTRY_POINT"); clientEntryPoint == "" {
		log.Fatalln("Environment value CLIENT_ENTRY_POINT is empty.")
	}
}
