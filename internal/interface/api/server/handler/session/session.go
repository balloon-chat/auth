package session

import (
	goenv "github.com/Netflix/go-env"
	"log"
)

const (
	sessionKey = "session"
)

var (
	ClientEntryPoint string
)

type Environment struct {
	ClientEntryPoint string `env:"CLIENT_ENTRY_POINT"`
}

func init() {
	var environment Environment
	_, err := goenv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables:", err)
	}

	if ClientEntryPoint = environment.ClientEntryPoint; ClientEntryPoint == "" {
		log.Fatalln("Environment variable CLIENT_ENTRY_POINT is empty.")
	}
}
