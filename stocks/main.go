package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/config"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/consumers"
)

func main() {

	config.InitConfig()

	go consumers.StockConsumer(config.KafkaStockConsumer)

	fmt.Println("Servidor Iniciando na porta 8083")
	log.Fatalln(http.ListenAndServe(":8083", nil))
}
