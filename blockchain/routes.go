package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SetupRoutes(blockchain *Blockchain) {
	http.HandleFunc("/transactions/new", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var tx Transaction
		if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
			http.Error(w, "Invalid transaction data", http.StatusBadRequest)
			return
		}

		// Add the transaction to the pool
		blockchain.TransactionPool = append(blockchain.TransactionPool, tx)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Transaction added to the pool")
	})

	http.HandleFunc("/mine", func(w http.ResponseWriter, r *http.Request) {
		// Mine a new block with transactions from the pool
		blockchain.AddBlock("New block mined")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "New block mined and added to the chain")
	})

	http.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		// Respond with the entire blockchain
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blockchain.Chain)
	})

	http.HandleFunc("/pool", func(w http.ResponseWriter, r *http.Request) {
		// Respond with the current transaction pool
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blockchain.TransactionPool)
	})
}
