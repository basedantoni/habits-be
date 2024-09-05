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
