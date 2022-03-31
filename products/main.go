package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Israel-Ferreira/techweek-hands-on/products/config"
	"github.com/Israel-Ferreira/techweek-hands-on/products/repositories"
	"github.com/Israel-Ferreira/techweek-hands-on/products/services"
	"github.com/Israel-Ferreira/techweek-hands-on/products/transport"
	"github.com/go-kit/log"
)

func main() {
	fmt.Println("Product Microsservice")

	logger := log.NewLogfmtLogger(os.Stderr)

	config.InitConfig()

	repo := repositories.NewRepository(config.Db)

	service := services.NewProductService(repo, config.Producer)

	httpServer := transport.NewHttpServer(service, logger)

	logger.Log("msg", "HTTP", "addr", "8082")
	logger.Log("err", http.ListenAndServe(":8082", httpServer))
}
