package google

import (
	"github.com/google/uuid"
	"net/http"
)

// Oauth GoogleのOAuth2による認証画面へリダイレクトする
func Oauth(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.FormValue("redirectUrl")

	//　認証完了時のリダイレクトURLを設定
	state := uuid.New().String()
	redirectUrls[state] = redirectUrl

	url := config.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}
