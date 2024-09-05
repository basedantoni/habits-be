// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: habits.sql

package database

import (
	"context"
	"database/sql"
)

const createHabit = `-- name: CreateHabit :one
INSERT INTO habits (
  id, title, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
)
RETURNING pk, id, title, created_at, updated_at, streak
`

type CreateHabitParams struct {
	ID        string
	Title     string
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}

func (q *Queries) CreateHabit(ctx context.Context, arg CreateHabitParams) (Habit, error) {
	row := q.db.QueryRowContext(ctx, createHabit,
		arg.ID,
		arg.Title,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Habit
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Streak,
	)
	return i, err
}

const deleteHabit = `-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = ?
`

func (q *Queries) DeleteHabit(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteHabit, id)
	return err
}

const getHabit = `-- name: GetHabit :one
SELECT pk, id, title, created_at, updated_at, streak FROM habits
WHERE id = ? LIMIT 1
`

func (q *Queries) GetHabit(ctx context.Context, id string) (Habit, error) {
	row := q.db.QueryRowContext(ctx, getHabit, id)
	var i Habit
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Streak,
	)
	return i, err
}

const listHabits = `-- name: ListHabits :many
SELECT pk, id, title, created_at, updated_at, streak FROM habits
ORDER BY title
LIMIT 20
`

func (q *Queries) ListHabits(ctx context.Context) ([]Habit, error) {
	rows, err := q.db.QueryContext(ctx, listHabits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Habit
	for rows.Next() {
		var i Habit
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Streak,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHabit = `-- name: UpdateHabit :exec
UPDATE habits
  SET title = ?, updated_at = ?, streak = ?
WHERE id = ?
`

type UpdateHabitParams struct {
	Title     string
	UpdatedAt sql.NullString
	Streak    sql.NullInt64
	ID        string
}

func (q *Queries) UpdateHabit(ctx context.Context, arg UpdateHabitParams) error {
	_, err := q.db.ExecContext(ctx, updateHabit,
		arg.Title,
		arg.UpdatedAt,
		arg.Streak,
		arg.ID,
	)
	return err
}
