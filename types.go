package main

import "basedantoni/habits-be/internal/database"

type apiConfig struct {
	DB *database.Queries
}

type Habit struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
