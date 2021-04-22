package google

import (
	"fmt"
	"github.com/Netflix/go-env"
	"github.com/balloon/auth/handler/oauth"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"log"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

// State OAuthで提供されるユーザーの識別子
type State = string

var (
	config       *oauth2.Config
	redirectUrls = map[State]string{}
	accessTokens = map[oauth.SessionID]string{}
)

type Environment struct {
	GoogleClientId     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}

	var environment Environment
	_, err = env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables:", err)
	}

	redirectUrl := fmt.Sprintf("%s/oauth/google/callback", oauth.BaseUrl)

	googleClientId := environment.GoogleClientId
	if googleClientId == "" {
		log.Fatalln("Environment variable GOOGLE_CLIENT_ID is empty.")
	}

	googleClientSecret := environment.GoogleClientSecret
	if googleClientSecret == "" {
		log.Fatalln("Environment variable GOOGLE_CLIENT_SECRET is empty.")
	}

	config = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
	}
}
