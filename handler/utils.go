package handler

import (
	"errors"
	"net/http"
	"strconv"
)

func ParseQuantity(r *http.Request) (int, error) {
	qtyStr := r.URL.Query().Get("quantity")
	if qtyStr == "" {
		return 0, errors.New("quantity query parameter is required")
	}
	quantity, err := strconv.Atoi(qtyStr)
	if err != nil {
		return 0, errors.New("quantity must be an integer")
	}
	if quantity < 1 {
		return 0, errors.New("quantity must be at least  1")
	}

	return quantity, nil

}
