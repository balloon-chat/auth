package cookie

import (
	goEnv "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
)

var (
	CookieDomain string
)

type Environment struct {
	CookieDomain string `env:"COOKIE_DOMAIN"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

	var environment Environment
	_, err = goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables:", err)
	}

	if CookieDomain = environment.CookieDomain; CookieDomain == "" {
		log.Fatalln("Environment variable COOKIE_DOMAIN is empty")
	}
}
