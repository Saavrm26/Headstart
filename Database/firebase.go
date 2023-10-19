package Database

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"hs/headstart/Utils"
	"log"
	"os"
)

var Fs *firestore.Client
var Bucket *storage.BucketHandle

func initializeFirestore(ctx context.Context, opts option.ClientOption) error {
	app, err := firebase.NewApp(ctx, nil, opts)

	if err != nil {
		return err
	}

	Fs, err = app.Firestore(ctx)

	if err != nil {
		return err
	}
	return nil
}

func initializeStorage(ctx context.Context, opts option.ClientOption) error {
	config := &firebase.Config{
		StorageBucket: os.Getenv("BUCKET_NAME") + ".appspot.com",
	}
	app, err := firebase.NewApp(context.Background(), config, opts)
	if err != nil {
		return err
	}
	St, err := app.Storage(ctx)
	if err != nil {
		return err
	}
	Bucket, err = St.DefaultBucket()
	return nil
}

func InitializeFirebase() error {
	// Use a service account
	ctx := context.Background()
	opts := option.WithCredentialsFile(Utils.FullPath("Secrets/firebase-service-account.json"))
	if err := initializeFirestore(ctx, opts); err != nil {
		return err
	}
	log.Println("Firestore successfully initialized")
	if err := initializeStorage(ctx, opts); err != nil {
		return err
	}
	log.Println("Bucket successfully initialized")

	return nil
}
