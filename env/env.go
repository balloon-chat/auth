package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DEBUG bool

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error while loading .env file")
	}
	DEBUG = os.Getenv("ENV") == "debug"
}
