package main

import (
	"FirstGo/handler"
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/api/v1/packDetails", handler.PackDetailsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf(":%s", port)

	fmt.Println("Listening on port :", port)
	http.ListenAndServe(address, nil)

}
