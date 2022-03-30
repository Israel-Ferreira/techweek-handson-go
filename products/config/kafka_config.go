package config

import (
	"fmt"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var Producer *kafka.Producer

func CreateProducer(bootstrapServer string) (*kafka.Producer, error) {
	config := &kafka.ConfigMap{"bootstrap.servers": bootstrapServer}

	fmt.Println(config)

	prdc, err := kafka.NewProducer(config)

	if err != nil {
		return nil, err
	}

	return prdc, nil
}

func SetProducer(bootstrapServer string) {
	prd, err := CreateProducer(bootstrapServer)

	if err != nil {
		log.Fatalln("erro ao conectar com o kafka")
	}

	Producer = prd
}
