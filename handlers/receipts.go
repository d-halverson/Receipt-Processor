package handlers

import (
	"encoding/json"
	"net/http"
	"receipts/models"
	"receipts/points"
	"receipts/storage"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handlers struct {
	storage *storage.ReceiptStorage
}

func NewHandlers(storage *storage.ReceiptStorage) *Handlers {
	return &Handlers{
		storage: storage,
	}
}

// Takes receipt in json format from request body.
func (h *Handlers) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := receipt.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New()
	h.storage.SetReceipt(id, &receipt)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Id{Id: id.String()})
}

// Calculates points for an existing receipt and returns them in response
func (h *Handlers) GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	parsedId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "receipt with id "+id+" not found", http.StatusNotFound)
		return
	}

	receipt := h.storage.GetReceipt(parsedId)
	if receipt == nil {
		http.Error(w, "receipt with id "+id+" not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.Points{Points: points.CalculatePoints(receipt)})
}
