package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) deleteHabitHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := cfg.DB.DeleteHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not index habits")
	}

	respondWithJSON(w, http.StatusNoContent, []string{})
}
