package handlers

import (
	"receipts/storage"

	"github.com/gorilla/mux"
)

/*
Creates a mux Router that has all of the server's api endpoint routing setup.

Created this function here instead of main package so that handler test files can use this
router as well as main file when program is ran.
*/
func CreateRouter() *mux.Router {
	receiptStorage := storage.NewReceiptStorage()
	handlers := NewHandlers(receiptStorage)
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")
	return router
}
