package kafkaprovider

import (
	"context"
	"example.com/m/provider"
	"github.com/sirupsen/logrus"

	"github.com/segmentio/kafka-go"
)

type KafkaProvider struct {
	chatWriter *kafka.Writer
}

func NewKafkaProvider() provider.KafkaProvider {

	kafkaHost := "localhost:9092"

	// chatWriter is a kafka writer for chat messages
	// Balancer:  &kafkaCustomChatBalancer{},

	chatWriter := &kafka.Writer{
		Addr:      kafka.TCP(kafkaHost),
		Topic:     "testing101",
		BatchSize: 1,
	}

	return &KafkaProvider{
		chatWriter: chatWriter,
	}
}

func (k *KafkaProvider) Publish(message []byte) {

	err := k.chatWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)

	if err != nil {
		logrus.Errorf("Publish: failed to write chat messages: %v", err)
		k.Reconnect()
	}
}

func (k *KafkaProvider) Reconnect() {
	kafkaHost := "localhost:9092"

	// chatWriter is a kafka writer for chat messages
	k.chatWriter = &kafka.Writer{
		Addr:  kafka.TCP(kafkaHost),
		Topic: "testing101",
	}
}

func (k *KafkaProvider) Close() {
	if err := k.chatWriter.Close(); err != nil {
		logrus.Errorf("error closing kafka chat connection")
	}
}
