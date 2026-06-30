package handler

import (
	"FirstGo/config"
	"FirstGo/service"
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

	packSize := config.GetAvailablePackSizes()

	if len(packSize) == 0 {
		http.Error(w, "Internal Server Error: Pack sizes are empty.", http.StatusInternalServerError)
		return
	}

	responseBody := service.CreatePackDetails(qty, packSize)

	json.NewEncoder(w).Encode(responseBody)

}
