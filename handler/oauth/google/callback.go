package google

import (
	"github.com/baloon/go/auth/app/infrastructure/firebase"
	"github.com/baloon/go/auth/handler/oauth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// OauthCallback GoogleのOauth認証のコールバックハンドラ
// トークンを取得し、セッションに登録する。
func OauthCallback(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")

	// トークンを取得
	code := c.Query("code")
	token, err := config.Exchange(c.Request.Context(), code)
	if err != nil {
		log.Println("error while generating token:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// セッションを作成
	store := oauth.Store
	session, _ := store.Get(c.Request, oauth.SessionCookieName)
	sessionId := uuid.New().String()
	session.Values[oauth.SessionIdCookieKey] = sessionId
	if err = session.Save(c.Request, c.Writer); err != nil {
		log.Println("error while writing cookie:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// アクセストークンを保持
	accessTokens[sessionId] = token.AccessToken

	profile, err := getUserProfile(token.AccessToken)
	if err != nil {
		return
	}

	found, _ := firebase.FindUserByEmail(c.Request.Context(), profile.Email)
	if !found {
		// 新規ユーザーの場合、sign inページへ
		c.Redirect(http.StatusFound, oauth.ClientSignInUrl)
		return
	}

	// ユーザーがすでに登録されている場合は、指定されているページへリダイレクトする
	state := c.Query("state")
	if redirectUrl, ok := redirectUrls[state]; ok {
		c.Redirect(http.StatusSeeOther, redirectUrl)
	} else {
		c.Redirect(http.StatusFound, oauth.ClientLoginUrl)
	}
}
