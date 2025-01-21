package api

import (
	"blockchain-voting/api/handler"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/vote", handler.VoteHandler) // Use the VoteHandler from the handler package
}
