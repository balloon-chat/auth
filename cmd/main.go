package main

import (
	"github.com/balloon/auth/internal/interface/api/server/handler/oauth/google"
	"github.com/balloon/auth/internal/interface/api/server/handler/session"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	g := r.Group("/oauth")
	{
		g = g.Group("/google")
		{
			g.GET("/", google.Oauth)
			g.GET("/callback", google.OauthCallback)
			g.GET("/result", google.GetOauthResult)
		}
	}

	g = r.Group("/session")
	{
		g.POST("/login", session.Login)
		g.OPTIONS("/login", setPostHeader)

		g.POST("/logout", session.Logout)
		g.OPTIONS("/logout", setPostHeader)

		g.GET("/profile", session.GetProfile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalln(err)
	}
}

func setPostHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", session.ClientEntryPoint)
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Status(http.StatusOK)
}
