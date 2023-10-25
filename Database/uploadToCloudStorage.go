package Database

import (
	"bytes"
	"context"
	"io"
	"time"
)

func UploadToCloudStorage(object string, buf *bytes.Buffer) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()
	wc := Bucket.Object(object).NewWriter(ctx)
	if _, err := io.Copy(wc, buf); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}
