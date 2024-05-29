package middlewares

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
)

func GoogleStorage() *GoogleCloudStorage {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &GoogleCloudStorage{
		Storage:    client,
		bucketName: os.Getenv("BUCKET_NAME"),
		projectID:  os.Getenv("GOOGLE_PROJECT_ID"),
		uploadPath: os.Getenv("BUCKET_PATH"),
	}
}

type GoogleCloudStorage struct {
	Storage    *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

func (gcs *GoogleCloudStorage) Path(path string) *GoogleCloudStorage {
	if len(path) != 0 {
		gcs.uploadPath = path
	}

	return gcs
}

func (gcs *GoogleCloudStorage) Upload(fileHeaders *multipart.FileHeader, file *multipart.File) (string, error) {

	ctx := context.Background()
	defer ctx.Done()

	path := fmt.Sprintf("%s/%s", gcs.uploadPath, fileHeaders.Filename)
	// Upload an object with storage.Writer.
	client := gcs.Storage.Bucket(gcs.bucketName)

	wc := client.Object(path).NewWriter(ctx)
	if _, err := io.Copy(wc, *file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	defer gcs.Storage.Close()

	link := fmt.Sprintf("https://storage.googleapis.com/nuria_bucket/%s", path)

	return link, nil
}
