package api

import "net/http"

func SetupRoutes() {
	http.HandleFunc("/vote", VoteHandler)
}
