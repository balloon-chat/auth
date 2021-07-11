package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/balloon/auth/internal/model"
	"sync"
)

var client *firestore.Client

type UserRepository interface {
	FindByLoginId(ctx context.Context, loginId string) (*model.User, error)
}

type firestoreUserRepository struct {
	sync.RWMutex
	client          *firestore.Client
	usersCollection *firestore.CollectionRef
}

func NewUserRepository(ctx context.Context) (UserRepository, error) {
	if client == nil {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firebase app: %v", err)
		}

		client, err = app.Firestore(ctx)
		if err != nil {
			return nil, fmt.Errorf("error while initializing firestore: %v", err)
		}
	}

	return &firestoreUserRepository{
		client:          client,
		usersCollection: client.Collection("users"),
	}, nil
}

func (repository *firestoreUserRepository) FindByLoginId(context context.Context, loginId string) (*model.User, error) {
	repository.RLock()
	defer repository.RUnlock()

	query := repository.usersCollection.Where("loginId", "==", loginId)
	docItr := query.Documents(context)

	doc, err := docItr.Next()
	if err != nil {
		return nil, fmt.Errorf("error while find user: %v", err)
	}

	var user model.User
	err = doc.DataTo(&user)
	if err != nil {
		return nil, fmt.Errorf("error while converting data: %v", err)
	}

	return &user, nil
}
