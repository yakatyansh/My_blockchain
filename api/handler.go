package api

import (
	"blockchain-voting/blockchain"
	"encoding/json"
	"net/http"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {

	voterID := "hashed_voter1"
	candidate := "Candidate A"

	newBlock := blockchain.GenerateBlock(blockchain.Blockchain[len(blockchain.Blockchain)-1], voterID, candidate)

	if blockchain.IsBlockValid(newBlock, blockchain.Blockchain[len(blockchain.Blockchain)-1]) {
		blockchain.Blockchain = append(blockchain.Blockchain, newBlock)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blockchain.Blockchain)
}
