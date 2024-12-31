package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize blockchain
	blockchain := NewBlockchain()
	TransactionPool = []Transaction{}

	// Set up routes
	SetupRoutes(blockchain)

	// Start the server
	port := ":8080"
	log.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
