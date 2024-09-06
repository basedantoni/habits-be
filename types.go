package main

import (
	"basedantoni/habits-be/internal/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

type apiConfig struct {
	DB *database.Queries
	Auth *oauth2.Config
}

type Claims struct {
	User database.User `json:"user"`
	jwt.StandardClaims
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type Habit struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	Streak int64 `json:"streak"`
}

type Contribution struct {
	Id        string `json:"id"`
	TimeSpent int64  `json:"timeSpent"`
	HabitId   int64 `json:"habitId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
