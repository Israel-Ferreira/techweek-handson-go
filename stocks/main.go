package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/config"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/consumers"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/repositories"
	"github.com/Israel-Ferreira/techweek-hands-on/stocks/src/services"
	routes "github.com/Israel-Ferreira/techweek-hands-on/stocks/src/transport"
	"github.com/go-kit/log"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	config.InitConfig()

	go consumers.StockConsumer(config.KafkaStockConsumer)

	stockRepo := repositories.NewStockRepository(config.Db)
	svc := services.NewStockService(stockRepo)

	server := routes.NewHttpServer(svc, logger)

	fmt.Println("Servidor Iniciando na porta 8083")

	logger.Log("err", http.ListenAndServe(":8083", server))
}
