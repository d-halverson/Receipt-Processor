package main

import (
	"fmt"
	"net/http"
	"receipts/handlers"
)

// Starts server listening on port 8080
func main() {
	router := handlers.CreateRouter()
	http.Handle("/", router)
	fmt.Println("Receipt Processor server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
