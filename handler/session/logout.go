package session

import (
	"github.com/baloon/go/auth/app/infrastructure/cookie"
	"github.com/baloon/go/auth/env"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Logout セッション情報を削除する
func Logout(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", ClientEntryPoint)
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Credentials", "true")

	c.SetCookie(
		sessionKey,
		"",
		-1, // Cookieを削除する
		"/",
		cookie.CookieDomain,
		!env.DEBUG,
		true,
	)
	c.Status(http.StatusOK)
}
