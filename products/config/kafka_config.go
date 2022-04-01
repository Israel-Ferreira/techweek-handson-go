package config

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

var Producer *kafka.Writer

func CreateProducer(bootstrapServer string) *kafka.Writer {
	fmt.Println(bootstrapServer)

	return &kafka.Writer{
		Addr:     kafka.TCP(bootstrapServer),
		Topic:    KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func SetProducer(bootstrapServer string) {
	Producer = CreateProducer(bootstrapServer)
}
