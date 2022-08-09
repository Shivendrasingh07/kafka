package provider

import (
	"context"
	"os"
	"time"
)

type StorageProvider interface {
	UploadV3(ctx context.Context, bucketName string, file *os.File, filePath string) error
	GetSharableURL(bucketName, fileName string, expireTimeInHours time.Duration) (string, error)
}

type KafkaProvider interface {
	Publish(message []byte)
	Reconnect()
	Close()
}

type WebSocketHubProvider interface {
	Run()
	Get() interface{}
	//Messengers() KafkaProvider
	//Stop()
	Subscribe()
	SubscribeAllPartitions()
	//SendOnlineStatusReport()
}
