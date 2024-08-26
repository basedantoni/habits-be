package main

import (
	"basedantoni/habits-be/internal/database"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func nullTimeToTimePtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func nullStringToTimePtr(ns sql.NullString) *time.Time {
	if ns.Valid {
		parsedTime, err := time.Parse(time.RFC3339, ns.String)
		if err != nil {
			return nil
		}
		
		return &parsedTime
	}
	return nil
}

func nullInt64ToInt64Ptr(ni sql.NullInt64) int64 {
	if ni.Valid {
		return ni.Int64
	}
	return 0
}

func databaseHabitToHabit(habit database.Habit) Habit {
	return Habit{
		Id:    habit.ID,
		Title: habit.Title,
		CreatedAt: nullStringToTimePtr(habit.CreatedAt),
		UpdatedAt: nullStringToTimePtr(habit.UpdatedAt),
	}
}


func databaseContributionToContribution(contribution database.Contribution) Contribution {
	return Contribution{
		Id:    contribution.ID,
		TimeSpent: contribution.TimeSpent,
		HabitId: nullInt64ToInt64Ptr(contribution.HabitID),
		CreatedAt: nullStringToTimePtr(contribution.CreatedAt),
		UpdatedAt: nullStringToTimePtr(contribution.UpdatedAt),
	}
}
