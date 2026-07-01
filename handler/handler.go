package handler

import (
	"FirstGo/config"
	"FirstGo/service"
	"context"
	"encoding/json"
	"net/http"
)

func PackDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	qty, err := ParseQuantity(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	packSizes := config.GetAvailablePackSizes(context.Background())
	if len(packSizes) == 0 {
		http.Error(w, "Internal Server Error: Pack sizes are empty.", http.StatusInternalServerError)
		return
	}

	responseBody := service.CreatePackDetails(qty, packSizes)

	json.NewEncoder(w).Encode(responseBody)

}
