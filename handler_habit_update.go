package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"basedantoni/habits-be/internal/database"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) updateHabitHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	type bodyParams struct {
		Title string `json:"title"`
	}

	params := bodyParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while decoding body params")
		return
	}

	err = cfg.DB.UpdateHabit(r.Context(), database.UpdateHabitParams{
		Title: params.Title,
		UpdatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not update new habit")
	}

	habit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find habit")
	}

	type response struct {
		Habit database.Habit `json:"habit"`
	}

	respondWithJSON(
		w, http.StatusOK,
		response{
			Habit: habit,
		},
	)
}
