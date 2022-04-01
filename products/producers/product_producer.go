package producers

import (
	"context"
	"encoding/json"

	"github.com/Israel-Ferreira/techweek-hands-on/products/data"
	"github.com/Israel-Ferreira/techweek-hands-on/products/models"
	"github.com/segmentio/kafka-go"
)

type ProductProducer struct {
	Kafka *kafka.Writer
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

	message := &kafka.Message{
		Key:   []byte(key["sku"]),
		Value: jsonResp,
	}

	if err = p.Kafka.WriteMessages(context.Background(), *message); err != nil {
		return err
	}

	return nil
}
