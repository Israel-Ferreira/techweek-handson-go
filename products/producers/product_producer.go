package producers

import (
	"encoding/json"

	"github.com/Israel-Ferreira/techweek-hands-on/products/config"
	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ProductProducer struct {
	Kafka *kafka.Producer
}

func (p ProductProducer) SendUpdateProductMsg(product models.Product) error {
	msg := data.EventProductMsg{Sku: product.Sku, Title: product.Title, Event: "UPDATE_PRODUCT"}

	if err := msg.IsValidEvent(); err != nil {
		return err
	}

	if err := p.sendMsg(msg); err != nil {
		return err
	}

	return nil

}

func (p ProductProducer) SendNewProductEventMsg(product models.Product) error {
	msg := data.EventProductMsg{Sku: product.Sku, Title: product.Title, Event: "NEW_PRODUCT"}

	if err := msg.IsValidEvent(); err != nil {
		return err
	}

	if err := p.sendMsg(msg); err != nil {
		return err
	}

	return nil
}

func (p ProductProducer) SendDeleteEventMsg(sku string) error {
	msg := data.EventProductMsg{Sku: sku, Title: "", Event: "DELETE_PRODUCT"}

	if err := msg.IsValidEvent(); err != nil {
		return err
	}

	if err := p.sendMsg(msg); err != nil {
		return err
	}

	return nil
}

func (p ProductProducer) sendMsg(msg data.EventProductMsg) error {
	jsonResp, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	key := map[string]string{
		"sku": msg.Sku,
	}

	topic := config.KafkaTopic

	message := &kafka.Message{
		Key:   []byte(key["sku"]),
		Value: jsonResp,
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
	}

	deliveryChan := make(chan kafka.Event, 10000)

	if err = p.Kafka.Produce(message, deliveryChan); err != nil {
		return err
	}

	return nil
}
