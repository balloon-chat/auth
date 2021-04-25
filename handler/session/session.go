package session

import (
	goenv "github.com/Netflix/go-env"
	"github.com/balloon/auth/env"
	"github.com/joho/godotenv"
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
	if env.DEBUG {
		_ = godotenv.Load(".env.develop")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

	var environment Environment
	_, err = goenv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables:", err)
	}

	if ClientEntryPoint = environment.ClientEntryPoint; ClientEntryPoint == "" {
		log.Fatalln("Environment variable CLIENT_ENTRY_POINT is empty.")
	}
}
