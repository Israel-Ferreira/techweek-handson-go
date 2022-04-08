package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/transport"
	"github.com/go-kit/log"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/config"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/consumers"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/repositories"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/services"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	config.InitConfig()

	repo := repositories.NewPriceRepo(config.Db)

	service := services.NewPriceService(repo)

	server := transport.NewServer(service, logger)

	go consumers.ProductConsumer(config.KafkaProductConsumer)

	fmt.Println("Servidor Iniciado na porta 8084")

	logger.Log("err", http.ListenAndServe(":8084", server))
}
