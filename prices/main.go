package main

import (
	"log"
	"net/http"
	"github.com/Israel-Ferreira/techweek-hands-on/prices/src/config"
)

func main() {

	config.InitConfig()

	log.Fatalln(http.ListenAndServe(":8084", nil))
}
