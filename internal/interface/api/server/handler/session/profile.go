package session

import (
	"context"
	firebase2 "github.com/balloon/auth/internal/infrastructure/firebase"
	"github.com/balloon/auth/internal/infrastructure/firestore"
	"github.com/balloon/auth/internal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	model.User
	// Firebaseログインで用いられるユーザー識別子
	LoginId string `json:"loginId"`
}

// GetProfile Firebaseによって作成されたセッション情報を用いて、ユーザーのプロフィールを取得
func GetProfile(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", ClientEntryPoint)
	c.Header("Access-Control-Allow-Credentials", "true")

	cookie, err := c.Cookie(sessionKey)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		log.Println("error while getting cookie:", err)
		return
	}

	client, err := firebase2.NewFirebaseAuthClient(context.Background())
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

	userRepository, err := firestore.NewUserRepository(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	user, err := userRepository.FindByLoginId(c.Request.Context(), decoded.UID)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, Response{
		User: model.User{
			ID:       user.ID,
			Name:     user.Name,
			PhotoUrl: user.PhotoUrl,
		},
		LoginId: decoded.UID,
	})
}
