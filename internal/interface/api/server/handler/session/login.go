package session

import (
	"context"
	"github.com/balloon/auth/env"
	"github.com/balloon/auth/internal/infrastructure/cookie"
	"github.com/balloon/auth/internal/infrastructure/firebase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Request struct {
	IdToken string `json:"idToken" bind:"required"`
}

// Login Firestoreが発行するトークンを用いてセッションを作成する
func Login(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", ClientEntryPoint)
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Credentials", "true")

	var request Request
	if err := c.BindJSON(&request); err != nil {
		log.Println("error while decoding request body:", err)
		c.Status(http.StatusBadRequest)
		return
	}

	// 有効期限: 5日
	expiresIn := 24 * time.Hour * 5

	client, err := firebase.NewFirebaseAuthClient(context.Background())
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	_, err = client.VerifyIDToken(c.Request.Context(), request.IdToken)
	if err != nil {
		log.Println("Invalid ID token", err)
		c.Status(http.StatusUnauthorized)
		return
	}

	// セッションCookieを作成
	ck, err := client.SessionCookie(c.Request.Context(), request.IdToken, expiresIn)
	if err != nil {
		log.Println("Failed to create session cookie:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.SetCookie(
		sessionKey,
		ck,
		int(expiresIn.Seconds()),
		"/",
		cookie.CookieDomain,
		!env.DEBUG,
		true,
	)
	c.Status(http.StatusOK)
}
