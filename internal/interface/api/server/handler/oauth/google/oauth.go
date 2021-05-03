package google

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// Oauth GoogleのOAuth2による認証画面へリダイレクトする
func Oauth(c *gin.Context) {
	redirectUrl := c.Query("redirectUrl")

	//　認証完了時のリダイレクトURLを設定
	state := uuid.New().String()
	redirectUrls[state] = redirectUrl

	url := config.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)
}
