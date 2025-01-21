package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CORS middleware to handle cross-origin requests
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allowed headers

		// Handle preflight (OPTIONS) requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// Handler to process voting requests
func VoteHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var vote struct {
		VoterID   string `json:"voterID"`
		Candidate string `json:"candidate"`
	}

	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log the vote to the console
	fmt.Printf("Received vote: %+v\n", vote)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Vote successfully recorded!",
	})
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()
	mux.Handle("/vote", http.HandlerFunc(VoteHandler))

	// Wrap the mux with the CORS middleware
	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", enableCORS(mux))
}
