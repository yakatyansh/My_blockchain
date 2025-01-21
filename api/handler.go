package api

import (
	"blockchain-voting/blockchain"
	"encoding/json"
	"net/http"
)

// VoteRequest represents the structure of the vote request sent from the frontend
type VoteRequest struct {
	VoterID   string `json:"voterID"`
	Candidate string `json:"candidate"`
}

// VoteResponse represents the structure of the response sent back to the frontend
type VoteResponse struct {
	Message    string             `json:"message"`
	Blockchain []blockchain.Block `json:"blockchain"`
}

// VoteHandler handles the vote submission
func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming JSON request
	var voteRequest VoteRequest
	err := json.NewDecoder(r.Body).Decode(&voteRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	previousBlock := blockchain.Blockchain[len(blockchain.Blockchain)-1]
	newBlock := blockchain.GenerateBlock(previousBlock, voteRequest.VoterID, voteRequest.Candidate)

	if blockchain.IsBlockValid(newBlock, previousBlock) {
		blockchain.Blockchain = append(blockchain.Blockchain, newBlock)

		// Send a success response back to the frontend
		response := VoteResponse{
			Message:    "Vote successfully recorded!",
			Blockchain: blockchain.Blockchain,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Send an error response if the block is invalid
		http.Error(w, "Invalid block, vote not recorded", http.StatusInternalServerError)
	}
}
