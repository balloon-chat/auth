package session

import (
	"context"
	"github.com/baloon/go/auth/app/infrastructure/firebase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	// Firebaseログインで用いられるユーザー識別子
	LoginId string `json:"loginId"`
}

// GetProfile Firebaseによって作成されたセッション情報を用いて、ユーザーのプロフィールを取得
func GetProfile(c *gin.Context) {
	cookie, err := c.Cookie(sessionKey)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		log.Println("error while getting cookie:", err)
		return
	}

	client, err := firebase.NewFirebaseAuthClient(context.Background())
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	decoded, err := client.VerifySessionCookie(c.Request.Context(), cookie)
	if err != nil {
		log.Println("error while verifying session cookie", err)
		c.Status(http.StatusUnauthorized)
		return
	}

	res := Response{
		LoginId: decoded.UID,
	}
	c.JSON(http.StatusOK, res)
}
