package Database

import (
	"cloud.google.com/go/storage"
	"fmt"
	"net/url"
	"os"
	"time"
)

func SignedURL(object string) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(144 * time.Hour), // expires 6 days from now
	}
	u, err := Bucket.SignedURL(object, opts)
	if err != nil {
		return "", err
	}
	return u, nil
}

func PublicURL(object string) string {
	bucketURL := os.Getenv("BUCKET_URL")
	apiURL := os.Getenv("FIREBASE_STORAGE_API")
	object = url.PathEscape(object)
	publicURL := fmt.Sprintf("%s/b/%s/o/%s?alt=media", apiURL, bucketURL, object)
	return publicURL
}
