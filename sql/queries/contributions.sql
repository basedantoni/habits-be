-- name: GetContribution :one
SELECT * FROM contributions
WHERE id = ? LIMIT 1;

-- name: ListContributions :many
SELECT * FROM contributions
LIMIT 20;

-- name: ListContributionsByHabit :many
SELECT * FROM contributions
WHERE habit_id = ?
LIMIT 20;

-- name: CreateContribution :one
INSERT INTO contributions (
  id, time_spent, habit_id, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateContribution :exec
UPDATE contributions
  SET time_spent = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteContribution :exec
DELETE FROM contributions
WHERE id = ?;
