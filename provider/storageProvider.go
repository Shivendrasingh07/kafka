package provider

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	storage2 "cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type GCPStorage struct {
	storageClient *storage2.Client
}

func NewStorageProvider(serviceKey string) StorageProvider {
	storageCl, err := storage2.NewClient(context.Background(), option.WithCredentialsJSON([]byte(serviceKey)))
	if err != nil {
		logrus.Fatalf("NewStorageProvider : %v", err)
	}

	return &GCPStorage{storageClient: storageCl}
}

func (gs GCPStorage) UploadV3(ctx context.Context, bucketName string, file *os.File, filePath string) error {
	bkt := gs.storageClient.Bucket(bucketName)

	obj := bkt.Object(filePath)

	attrs, err := bkt.Attrs(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	fmt.Printf("bucket %s, created at %s, is located in %s with storage class %s\n",
		attrs.Name, attrs.Created, attrs.Location, attrs.StorageClass)

	// Write something to obj.
	// w implements io.Writer.
	w := obj.NewWriter(ctx)

	if _, err := io.Copy(w, file); err != nil {
		// TODO: Handle error.
	}

	// Close, just like writing a file.
	if err := w.Close(); err != nil {
		// TODO: Handle error.
	}

	// Read it back.
	r, err := obj.NewReader(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer r.Close()

	err = w.Close()
	if err != nil {
		logrus.Errorf("This object contains text. %v \n", err)
	}

	// Prints "This object contains text."

	//_, err = obj.Update(ctx, storage2.ObjectAttrsToUpdate{
	//	Metadata:    metadata,
	//	ContentType: contentType,
	//})

	return nil
}

func (gs GCPStorage) GetSharableURL(bucketName, fileName string, expireTimeInHours time.Duration) (string, error) {
	jwt, err := google.JWTConfigFromJSON([]byte(os.Getenv("GCP_Key")))
	if err != nil {
		return "", err
	}

	opts := &storage2.SignedURLOptions{
		GoogleAccessID: jwt.Email,
		PrivateKey:     jwt.PrivateKey,
		Method:         "GET",
		Expires:        time.Now().Add(expireTimeInHours * time.Hour),
	}

	url, err := storage2.SignedURL(bucketName, fileName, opts)

	if err != nil {
		return "", err
	}

	return url, nil
}
