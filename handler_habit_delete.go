package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (cfg *apiConfig) deleteHabitHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Sprintf("Deleting habit with id %s", id)	
	err := cfg.DB.DeleteHabit(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusNoContent, []string{})
}
