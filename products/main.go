package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Product Microsservice")
	
	log.Fatalln(http.ListenAndServe(":8082", nil))
}
