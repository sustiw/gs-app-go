package main

import (
	"FirstGo/config"
	"FirstGo/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/packDetails", handler.PackDetailsHandler)
	address := config.GetServerPort()

	fmt.Printf("Elastic Beanstalk: Launching Go Application on interface address %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
