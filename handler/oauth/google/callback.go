package google

import (
	"github.com/baloon/go/auth/handler/oauth"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// OauthCallback GoogleのOauth認証のコールバックハンドラ
// トークンを取得し、セッションに登録する。
func OauthCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// トークンを取得
	code := r.FormValue("code")
	token, err := config.Exchange(r.Context(), code)
	if err != nil {
		log.Println("error while generating token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// セッションを作成
	store := oauth.Store
	session, _ := store.Get(r, oauth.SessionCookieName)
	sessionId := uuid.New().String()
	session.Values[oauth.SessionIdCookieKey] = sessionId
	if err = session.Save(r, w); err != nil {
		log.Println("error while writing cookie:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// アクセストークンを保持
	accessTokens[sessionId] = token.AccessToken

	// リダイレクト
	state := r.FormValue("state")
	if redirectUrl, ok := redirectUrls[state]; ok {
		// ユーザーがすでに登録されている場合は、指定されているページへリダイレクトする
		http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
