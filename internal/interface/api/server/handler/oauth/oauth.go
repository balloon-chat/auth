package oauth

import (
	goEnv "github.com/Netflix/go-env"
	"github.com/balloon/auth/env"
	"github.com/balloon/auth/internal/infrastructure/cookie"
	"github.com/gorilla/sessions"
	"log"
	"time"
)

// SessionID Cookieに保存するセッションのID
type SessionID = string

const (
	SessionCookieName  = "session"
	SessionIdCookieKey = "session_id"
)

var (
	ClientSignInUrl string
	ClientLoginUrl  string
	BaseUrl         string
	Store           *sessions.CookieStore
)

type Environment struct {
	BaseUrl          string `env:"BASE_URL"`
	ClientSignInUrl  string `env:"CLIENT_SIGN_IN_URL"`
	ClientLoginUrl   string `env:"CLIENT_LOGIN_URL"`
	SessionSecretKey string `env:"SESSION_SECRET_KEY"`
}

func init() {
	env.LoadEnv()

	var environment Environment
	_, err := goEnv.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatalln("error while parsing environment variables:", err)
	}

	ClientSignInUrl = environment.ClientSignInUrl
	if ClientSignInUrl == "" {
		log.Fatalln("Environment variable CLIENT_SIGN_IN_URL is empty")
	}

	ClientLoginUrl = environment.ClientLoginUrl
	if ClientLoginUrl == "" {
		log.Fatalln("Environment variable CLIENT_LOGIN_URL is empty")
	}

	BaseUrl = environment.BaseUrl
	if BaseUrl == "" {
		log.Fatalln("Environment variable BASE_URL is empty")
	}

	sessionSecretKey := environment.SessionSecretKey
	if sessionSecretKey == "" {
		log.Fatalln("Environment variable SESSION_SECRET_KEY is empty")
	}

	expiresIn := 24 * time.Hour * 5

	Store = sessions.NewCookieStore([]byte(sessionSecretKey))
	Store.Options = &sessions.Options{
		HttpOnly: true,
		Path:     "/",
		Domain:   cookie.CookieDomain,
		Secure:   !env.DEBUG,
		MaxAge:   int(expiresIn.Seconds()),
	}
}
