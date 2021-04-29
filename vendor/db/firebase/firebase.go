package firebase

import (
	"context"
	"huntsub/huntsub-map-server/x/mlog"
	"log"

	"cloud.google.com/go/firestore"
	store "firebase.google.com/go"
	"google.golang.org/api/option"
)

var firebaseLog = mlog.NewTagLog("firebaseLog")

type Firebase struct {
	Ctx         context.Context
	ClientStore *firestore.Client
}

func NewClientFirebase() (*Firebase, error) {
	var _firebase = &Firebase{}
	_firebase.Ctx = context.Background()

	sa := option.WithCredentialsFile("serviceAccount.json")
	app, err := store.NewApp(_firebase.Ctx, nil, sa)
	if err != nil {
		firebaseLog.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	_firebase.ClientStore, err = app.Firestore(_firebase.Ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return _firebase, nil
}
