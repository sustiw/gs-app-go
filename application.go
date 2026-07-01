package main

import (
	"FirstGo/config" // Import your config package to pull the port utility
	"FirstGo/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Register your API endpoint routing handler
	http.HandleFunc("/api/v1/packDetails", handler.PackDetailsHandler)

	// 2. Fetch the correctly formatted server address string (e.g. ":5000")
	address := config.GetServerPort()

	// 3. Print out a clear initialization message to the log engine
	fmt.Printf("Elastic Beanstalk: Launching Go Application on interface address %s\n", address)

	// 4. Start the listener globally so the AWS Nginx proxy can route inbound traffic
	log.Fatal(http.ListenAndServe(address, nil))
}
