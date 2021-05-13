package main

import (
	"github.com/balloon/auth/env"
	"github.com/balloon/auth/internal/interface/api/server"
	"log"
	"os"
)

func main() {
	env.LoadEnv()
	log.Println(os.Getenv("CLIENT_LOGIN_URL"))
	server.ServeHTTP(false)
}
