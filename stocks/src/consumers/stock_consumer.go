package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/config"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/data"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/models"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/repositories"
	"github.com/segmentio/kafka-go"
)

func StockConsumer(reader *kafka.Reader) {

	repo := repositories.NewStockRepository(config.Db)

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

				stockItem := models.Stock{
					Sku:   eventMsg.Sku,
					Title: eventMsg.Title,
				}

				if err = repo.CreateStockItem(stockItem); err != nil {
					log.Println(err)
				}

			} else if eventMsg.Event == "DELETE_PRODUCT" {
				if err = repo.DeleteStockItemBySku(eventMsg.Sku); err != nil {
					log.Println(err)
				}
			}
		}

		fmt.Println(eventMsg)
	}
}
