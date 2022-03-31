package consumers

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StockConsumer(reader *kafka.Reader) {

	for {
		msg, err := reader.ReadMessage(context.Background())

		if err != nil {
			log.Println("Erro ao ler a mensagem do topico")
		}

		fmt.Println(string(msg.Value))
	}
}
