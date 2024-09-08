package main

import (
	"basedantoni/habits-be/internal/database"
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func (cfg *apiConfig) googleAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	state := r.URL.Query().Get("state")
    originalDestination, err := base64.StdEncoding.DecodeString(state)
    if err != nil || len(originalDestination) == 0 {
        originalDestination = []byte("/")
    }
	
	// Get the authorization code from the URL query parameters
	code := r.URL.Query().Get("code")
	
	if code == "" {
		http.Error(w, "No code in request", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	tok, err := cfg.Auth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to get the user's info
	client := cfg.Auth.Client(context.Background(), tok)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := cfg.DB.GetUser(r.Context(), userInfo.Email)
	if err != nil {
		user, err = cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID: nanoid.New(),
			Email: userInfo.Email,
			Password: sql.NullString{},
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Could not create new user")
		}
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	var jwtKey = []byte(os.Getenv("SESSION_KEY"))
	if jwtKey == nil {
		respondWithError(w, http.StatusBadRequest, "No session key variable")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to frontend with token as query parameter
	redirectURL := fmt.Sprintf(
		"%s?token=%s&redirect=%s", 
		os.Getenv("CLIENT_BASE_URL"), 
		tokenString, 
		url.QueryEscape(string(originalDestination)),
	)
	
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}