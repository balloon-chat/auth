package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DEBUG bool
var VERSION string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error while loading .env file:", err)
	}
	DEBUG = os.Getenv("ENV") == "development"
	VERSION = os.Getenv("VERSION")
}

func LoadEnv() {
	err := godotenv.Load(fmt.Sprintf(".env.%s", VERSION))
	if err != nil {
		log.Fatalln("error while loading .env file:", err)
	}
}
