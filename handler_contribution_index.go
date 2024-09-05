package main

import (
	"basedantoni/habits-be/internal/database"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) indexContributionHandler(w http.ResponseWriter, r *http.Request) {
	databaseContributions, err := cfg.DB.ListContributions(r.Context())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index contributions")
	}

	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}

func (cfg *apiConfig) indexHabitContributionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	year := r.URL.Query().Get("year")

	habit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find contributions")
	}

	databaseContributions := []database.Contribution{}

	if year != "" {
		databaseContributions, err = cfg.DB.GetContributionsByYear(
			r.Context(),
			database.GetContributionsByYearParams{
				HabitID: sql.NullInt64{Int64: habit.Pk, Valid: true},
				CreatedAt: sql.NullString{String: year, Valid: true},
			},
		)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Could not index contributions")
		}
	} else {
		databaseContributions, err = cfg.DB.GetContributionsByPastYear(r.Context(), sql.NullInt64{Int64: habit.Pk, Valid: true})
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Could not index contributions")
		}
	}


	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}

func (cfg *apiConfig) indexHabitContributionByYearHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	year := chi.URLParam(r, "year")

	habit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find contributions")
	}

	databaseContributions, err := cfg.DB.GetContributionsByYear(
		r.Context(),
		database.GetContributionsByYearParams{
			HabitID: sql.NullInt64{Int64: habit.Pk, Valid: true},
			CreatedAt: sql.NullString{String: year, Valid: true},
		},
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index contributions")
	}

	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}

func (cfg *apiConfig) indexHabitContributionByPastYearHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	habit, err := cfg.DB.GetHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find contributions")
	}

	databaseContributions, err := cfg.DB.GetContributionsByPastYear(
		r.Context(),
		sql.NullInt64{Int64: habit.Pk, Valid: true},
	)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index contributions")
	}

	contributions := []Contribution{}
	for _, c := range databaseContributions {
		contributions = append(contributions, databaseContributionToContribution(c))
	}

	respondWithJSON(w, http.StatusOK, contributions)
}