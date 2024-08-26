package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) indexContributionHandler(w http.ResponseWriter, r *http.Request) {
	databaseContributions, err := cfg.DB.ListContributions(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index habits")
	}

	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}

func (cfg *apiConfig) indexHabitContributionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	habit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find habit")
	}

	databaseContributions, err := cfg.DB.ListContributionsByHabit(r.Context(), sql.NullInt64{Int64: habit.Pk, Valid: true})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index habits")
	}

	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}