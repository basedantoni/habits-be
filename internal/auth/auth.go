package auth

import (
	"basedantoni/habits-be/internal/database"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Claims struct {
	User database.User `json:"user"`
	jwt.StandardClaims
}

func Authenticate(next http.Handler) http.Handler {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var jwtKey = []byte(os.Getenv("SESSION_KEY"))
		if jwtKey == nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", claims.User.Pk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}