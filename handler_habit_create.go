package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"basedantoni/habits-be/internal/database"

	"github.com/aidarkhanov/nanoid"
)

func (cfg *apiConfig) createHabitHandler(w http.ResponseWriter, r *http.Request) {
	type bodyParams struct {
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
		UserID: sql.NullInt64{
			Valid: r.Context().Value("userID").(int64) != 0,
			Int64: r.Context().Value("userID").(int64),
		},
		CreatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
		UpdatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	type response struct {
		Habit Habit `json:"habit"`
	}

	respondWithJSON(
		w, http.StatusOK,
		response{
			Habit: databaseHabitToHabit(habit),
		},
	)
}
