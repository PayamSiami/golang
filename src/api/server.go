package api

import (
	"api/router"
	"auto"
	"config"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListing [::]:%d\n", config.PORT)
	Listen(config.PORT)
}

func Listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
