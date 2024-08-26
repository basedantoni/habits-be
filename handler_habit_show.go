package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) showHabitHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	dbHabit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find habit")
	}

	habit := databaseHabitToHabit(dbHabit)

	respondWithJSON(w, http.StatusOK, habit)
}