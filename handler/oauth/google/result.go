package google

import (
	"encoding/json"
	"fmt"
	"github.com/baloon/go/auth/app/infrastructure/firebase"
	"github.com/baloon/go/auth/handler/oauth"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	AccessToken string `json:"accessToken"`
	Name        string `json:"name"`
	PhotoUrl    string `json:"photoUrl"`
	// NewUser 新規ユーザーかどうか
	NewUser bool `json:"newUser"`
}

// GetOauthResult OAuth認証によって取得したアクセストークンを用いて、ユーザーのプロフィールを取得する。
func GetOauthResult(w http.ResponseWriter, r *http.Request) {
	store := oauth.Store
	session, _ := store.Get(r, oauth.SessionCookieName)
	sessionId, ok := session.Values[oauth.SessionIdCookieKey]
	if !ok {
		return
	}

	switch st := sessionId.(type) {
	case string:
		// アクセストークンから、ユーザープロファイルを取得
		accessToken, ok := accessTokens[st]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		profile, err := getUserProfile(accessToken)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// 登録状態を確認
		found, _ := firebase.FindUserByEmail(r.Context(), profile.Email)

		res := Response{
			AccessToken: accessToken,
			Name:        profile.Name,
			PhotoUrl:    profile.PhotoUrl,
			NewUser:     !found,
		}
		if err = json.NewEncoder(w).Encode(res); err != nil {
			log.Println("error while encoding response:", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusUnauthorized)
	}
}

type UserProfile struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"picture"`
}

func getUserProfile(accessToken string) (*UserProfile, error) {
	userInfoUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", url.QueryEscape(accessToken))
	res, err := http.Get(userInfoUrl)
	if err != nil {
		return nil, fmt.Errorf("error while getting user email: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid access token")
	}

	var profile UserProfile
	if err = json.NewDecoder(res.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("error while decoding response body: %v", err)
	}

	return &profile, nil
}
