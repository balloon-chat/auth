package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/balloon/auth/env"
	"github.com/joho/godotenv"
	"log"
)

var client *auth.Client

func init() {
	if env.DEBUG {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("error while loading .env file:", err)
		}
	}

	c, err := NewFirebaseAuthClient(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	client = c
}

func NewFirebaseAuthClient(ctx context.Context) (*auth.Client, error) {
	if client == nil {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firebase app: %v", err)
		}

		client, err = app.Auth(ctx)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firebase auth client: %v", err)
		}
	}

	return client, nil
}

func FindUserByEmail(context context.Context, mailAddress string) (found bool, err error) {
	_, err = client.GetUserByEmail(context, mailAddress)
	if err != nil {
		return false, fmt.Errorf("error while getting user: %v", err)
	}
	return true, nil
}
