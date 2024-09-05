// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: contributions.sql

package database

import (
	"context"
	"database/sql"
)

const createContribution = `-- name: CreateContribution :one
INSERT INTO contributions (
  id, time_spent, habit_id, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING pk, id, time_spent, habit_id, created_at, updated_at
`

type CreateContributionParams struct {
	ID        string
	TimeSpent int64
	HabitID   sql.NullInt64
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}

func (q *Queries) CreateContribution(ctx context.Context, arg CreateContributionParams) (Contribution, error) {
	row := q.db.QueryRowContext(ctx, createContribution,
		arg.ID,
		arg.TimeSpent,
		arg.HabitID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Contribution
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.TimeSpent,
		&i.HabitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteContribution = `-- name: DeleteContribution :exec
DELETE FROM contributions
WHERE id = ?
`

func (q *Queries) DeleteContribution(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteContribution, id)
	return err
}

const getContribution = `-- name: GetContribution :one
SELECT pk, id, time_spent, habit_id, created_at, updated_at FROM contributions
WHERE id = ? LIMIT 1
`

func (q *Queries) GetContribution(ctx context.Context, id string) (Contribution, error) {
	row := q.db.QueryRowContext(ctx, getContribution, id)
	var i Contribution
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.TimeSpent,
		&i.HabitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getContributionsByPastYear = `-- name: GetContributionsByPastYear :many
SELECT pk, id, time_spent, habit_id, created_at, updated_at
FROM contributions
WHERE habit_id = ?
AND created_at >= date('now', '-1 year')
ORDER BY created_at ASC
`

func (q *Queries) GetContributionsByPastYear(ctx context.Context, habitID sql.NullInt64) ([]Contribution, error) {
	rows, err := q.db.QueryContext(ctx, getContributionsByPastYear, habitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contribution
	for rows.Next() {
		var i Contribution
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.TimeSpent,
			&i.HabitID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getContributionsByYear = `-- name: GetContributionsByYear :many
SELECT pk, id, time_spent, habit_id, created_at, updated_at FROM contributions
WHERE habit_id = ? AND strftime('%Y', created_at) = ?
ORDER BY created_at ASC
`

type GetContributionsByYearParams struct {
	HabitID   sql.NullInt64
	CreatedAt sql.NullString
}

func (q *Queries) GetContributionsByYear(ctx context.Context, arg GetContributionsByYearParams) ([]Contribution, error) {
	rows, err := q.db.QueryContext(ctx, getContributionsByYear, arg.HabitID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contribution
	for rows.Next() {
		var i Contribution
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.TimeSpent,
			&i.HabitID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getLastContribution = `-- name: GetLastContribution :one
SELECT pk, id, time_spent, habit_id, created_at, updated_at FROM contributions
WHERE habit_id = ?
ORDER BY created_at DESC
LIMIT 1
`

func (q *Queries) GetLastContribution(ctx context.Context, habitID sql.NullInt64) (Contribution, error) {
	row := q.db.QueryRowContext(ctx, getLastContribution, habitID)
	var i Contribution
	err := row.Scan(
		&i.Pk,
		&i.ID,
		&i.TimeSpent,
		&i.HabitID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listContributions = `-- name: ListContributions :many
SELECT pk, id, time_spent, habit_id, created_at, updated_at FROM contributions
LIMIT 20
`

func (q *Queries) ListContributions(ctx context.Context) ([]Contribution, error) {
	rows, err := q.db.QueryContext(ctx, listContributions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contribution
	for rows.Next() {
		var i Contribution
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.TimeSpent,
			&i.HabitID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listContributionsByHabit = `-- name: ListContributionsByHabit :many
SELECT pk, id, time_spent, habit_id, created_at, updated_at FROM contributions
WHERE habit_id = ?
LIMIT 20
`

func (q *Queries) ListContributionsByHabit(ctx context.Context, habitID sql.NullInt64) ([]Contribution, error) {
	rows, err := q.db.QueryContext(ctx, listContributionsByHabit, habitID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Contribution
	for rows.Next() {
		var i Contribution
		if err := rows.Scan(
			&i.Pk,
			&i.ID,
			&i.TimeSpent,
			&i.HabitID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateContribution = `-- name: UpdateContribution :exec
UPDATE contributions
  SET time_spent = ?, updated_at = ?
WHERE id = ?
`

type UpdateContributionParams struct {
	TimeSpent int64
	UpdatedAt sql.NullString
	ID        string
}

func (q *Queries) UpdateContribution(ctx context.Context, arg UpdateContributionParams) error {
	_, err := q.db.ExecContext(ctx, updateContribution, arg.TimeSpent, arg.UpdatedAt, arg.ID)
	return err
}
