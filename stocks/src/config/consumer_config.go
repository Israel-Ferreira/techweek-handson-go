package config

import "github.com/segmentio/kafka-go"

func ConsumerConfig(bootstrapServer, topic, groupId string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{bootstrapServer},
		Topic:   topic,
		GroupID: groupId,
	})
}
