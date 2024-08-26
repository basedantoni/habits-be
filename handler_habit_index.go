package main

import (
	"net/http"
)

func (cfg *apiConfig) indexHabitHandler(w http.ResponseWriter, r *http.Request) {
	databaseHabits, err := cfg.DB.ListHabits(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index habits")
	}

	habits := []Habit{}
	for _, h := range databaseHabits {
		habits = append(habits, databaseHabitToHabit(h))
	}

	respondWithJSON(w, http.StatusOK, habits)
}
