// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
)

type Contribution struct {
	Pk        int64
	ID        string
	TimeSpent int64
	HabitID   sql.NullInt64
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}

type Habit struct {
	Pk        int64
	ID        string
	Title     string
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
	Streak    sql.NullInt64
}
