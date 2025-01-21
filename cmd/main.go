package main

import (
	"blockchain-voting/api"
	"blockchain-voting/blockchain"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Initialize the blockchain with the genesis block
	mux := http.NewServeMux()
	blockchain.InitializeBlockChain()

	// Setup the API routes
	api.SetupRoutes(mux)

	// Start the server
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
