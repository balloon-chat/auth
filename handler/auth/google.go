package auth

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

const (
	authorizeEndpoint    = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint        = "https://www.googleapis.com/oauth2/v4/token"
	sessionCookieKey     = "session"
	accessTokenCookieKey = "access_token"
)

var (
	config       *oauth2.Config
	redirectUrls = map[string]string{}
	store        *sessions.CookieStore
)

func init() {
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if googleClientId == "" {
		log.Fatalln("Environment value GOOGLE_CLIENT_ID is not specified.")
	}

	if googleClientSecret == "" {
		log.Fatalln("Environment value GOOGLE_CLIENT_SECRET is not specified.")
	}

	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		log.Fatalln("Environment value BASE_URL is not specified")
	}
	redirectUrl := fmt.Sprintf("%s/oauth/google/callback", baseUrl)

	config = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       []string{"profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
	}

	store = sessions.NewCookieStore([]byte("secret"))
	store.Options.HttpOnly = true
}

func OauthGoogle(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.FormValue("redirectUrl")

	u, err := uuid.NewUUID()
	if err != nil {
		log.Println("error while generating new uuid:", err)
		return
	}

	state := u.String()
	redirectUrls[state] = redirectUrl
	url := config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

func OauthCallbackGoogle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	code := r.FormValue("code")
	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		log.Println("error while generating token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session, _ := store.Get(r, sessionCookieKey)

	state := r.FormValue("state")
	if redirectUrl, ok := redirectUrls[state]; ok {
		session.Values[accessTokenCookieKey] = token.AccessToken

		err = session.Save(r, w)
		if err != nil {
			log.Println("error while writing cookie:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, sessionCookieKey)
	accessToken, ok := session.Values[accessTokenCookieKey]
	if !ok {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch at := accessToken.(type) {
	case string:
		res := struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: at,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Println("error while encoding response:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
