package firebase

// import (
// 	"context"
// 	"log"

// 	store "firebase.google.com/go"
// 	"firebase.google.com/go/v4/storage"
// 	"google.golang.org/api/option"
// )

// type FireBaseStore struct {
// 	Ctx           context.Context
// 	ClientStorage *storage.Client
// }

// func NewClientStorage() (*FireBaseStore, error) {
// 	var _firebase = &FireBaseStore{}
// 	_firebase.Ctx = context.Background()
// 	config := &store.Config{
// 		StorageBucket: "huntsub-messager-4e58f.appspot.com",
// 	}
// 	sa := option.WithCredentialsFile("serviceAccount.json")
// 	app, err := store.NewApp(_firebase.Ctx, config, sa)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_firebase.ClientStorage, err = app.Storage(_firebase.Ctx)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// bucket, err := client.DefaultBucket()
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	return _firebase, nil
// }
