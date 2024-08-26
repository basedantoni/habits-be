package main

import (
	"basedantoni/habits-be/internal/database"
	"time"
)

type apiConfig struct {
	DB *database.Queries
}

type Habit struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Contribution struct {
	Id        string `json:"id"`
	TimeSpent int64  `json:"time_spent"`
	HabitId   int64 `json:"habit_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
