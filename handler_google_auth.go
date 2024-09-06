package main

import (
	"net/http"

	"golang.org/x/oauth2"
)

func (cfg *apiConfig) googleAuthHandler(w http.ResponseWriter, r *http.Request) {
	url := cfg.Auth.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}