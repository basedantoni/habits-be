package main

import (
	"basedantoni/habits-be/internal/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/aidarkhanov/nanoid"
)

func (cfg *apiConfig) createContributionHandler(w http.ResponseWriter, r *http.Request) {
	type bodyParams struct {
		HabitId string `json:"habit_id"`
	}

	params := bodyParams{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while decoding body params")
		return
	}

	dbHabit, err := cfg.DB.GetHabit(r.Context(), params.HabitId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find habit")
	}

	contribution, err := cfg.DB.CreateContribution(r.Context(), database.CreateContributionParams{
		ID: nanoid.New(),
		HabitID: sql.NullInt64{Int64: dbHabit.Pk, Valid: true},
		TimeSpent: 0,
		CreatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
		UpdatedAt: sql.NullString{String: time.Now().Format(time.RFC3339), Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not create new contribution")
	}

	type response struct {
		Contribution database.Contribution `json:"contribution"`
	}

	respondWithJSON(
		w, http.StatusOK,
		response{
			Contribution: contribution,
		},
	)

}