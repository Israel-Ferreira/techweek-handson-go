package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/config"
)

func main() {

	config.InitConfig()

	fmt.Println("Servidor Iniciado na porta 8084")
	log.Fatalln(http.ListenAndServe(":8084", nil))
}
