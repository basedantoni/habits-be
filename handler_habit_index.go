package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) indexHabitHandler(w http.ResponseWriter, r *http.Request) {
	databaseHabits, err := cfg.DB.ListHabits(r.Context())
	if err != nil {
		e := fmt.Sprintf("%s %s", "Could not list habits", err.Error())
		respondWithError(w, http.StatusBadRequest, e)
	}

	habits := []Habit{}
	for _, h := range databaseHabits {
		habits = append(habits, databaseHabitToHabit(h))
	}

	respondWithJSON(w, http.StatusOK, habits)
}
