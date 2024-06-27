package main

import (
	"fmt"
	"net/http"
	"receipts/handlers"

	"github.com/gorilla/mux"
)

// Starts server listening on port 8080
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Receipt Processor server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
