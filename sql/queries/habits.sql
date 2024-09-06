-- name: GetHabit :one
SELECT * FROM habits
WHERE id = ? LIMIT 1;

-- name: ListHabits :many
SELECT * FROM habits
WHERE user_id = ?
LIMIT 20;

-- name: CreateHabit :one
INSERT INTO habits (
  id, title, user_id, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateHabit :exec
UPDATE habits
  SET title = ?, updated_at = ?, streak = ?
WHERE id = ?;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = ?;
