package main

import (
	"encoding/json"
	"net/http"

	"basedantoni/habits-be/internal/database"

	"github.com/aidarkhanov/nanoid"
)

func (cfg *apiConfig) createHabitHandler(w http.ResponseWriter, r *http.Request) {
	type bodyParams struct {
		Id    string `json:"id"`
		Title string `json:"title"`
	}

	params := bodyParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while decoding body params")
		return
	}

	habit, err := cfg.DB.CreateHabit(r.Context(), database.CreateHabitParams{
		ID:    nanoid.New(),
		Title: params.Title,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not create new feed")
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
