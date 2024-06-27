package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipts/models"
)

// Takes receipt in json format from request body.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {

	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(receipt.PurchaseDate)
	fmt.Println(receipt.PurchaseTime)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Looks good!"))
}
