package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

func (cfg *apiConfig) googleAuthHandler(w http.ResponseWriter, r *http.Request) {
    originalDestination := r.URL.Query().Get("redirect")
    if originalDestination == "" {
        originalDestination = "/"
    }

	fmt.Println(originalDestination)
    state := base64.StdEncoding.EncodeToString([]byte(originalDestination))
    url := cfg.Auth.AuthCodeURL(state, oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}