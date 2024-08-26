package main

import (
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, response{Status: "ok"})
}
