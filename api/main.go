package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	config.LoadConfig()
}

func main() {
	r := router.GenerateRouter()
	fmt.Printf("Listening on port %d...", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
