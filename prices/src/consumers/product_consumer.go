package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/config"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/models"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/repositories"
	"github.com/segmentio/kafka-go"
)

func ProductConsumer(reader *kafka.Reader) {

	repo := repositories.NewPriceRepo(config.Db)

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("Erro ao ler a mensagem do topico")
		}

		var eventMsg data.EventMsg

		if err = json.Unmarshal(msg.Value, &eventMsg); err != nil {
			log.Println(err)
		} else {

			if eventMsg.Event == "NEW_PRODUCT" {

				productItem := models.Price{
					Title: eventMsg.Title,
					Sku:   eventMsg.Sku,
				}

				if _, err = repo.AddItem(productItem); err != nil {
					log.Println(err)
				}

			} else if eventMsg.Event == "DELETE_PRODUCT" {
				if err = repo.DeleteBySku(eventMsg.Sku); err != nil {
					log.Println(err)
				}
			}
		}

		fmt.Println(eventMsg)
	}
}
