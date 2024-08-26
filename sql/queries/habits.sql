-- name: GetHabit :one
SELECT * FROM habits
WHERE id = ? LIMIT 1;

-- name: ListHabits :many
SELECT * FROM habits
ORDER BY title;

-- name: CreateHabit :one
INSERT INTO habits (
  id, title
) VALUES (
  ?, ?
)
RETURNING *;

-- name: UpdateHabit :exec
UPDATE habits
  SET title = ?
WHERE id = ?;

-- name: DeleteHabit :exec
DELETE FROM habits
WHERE id = ?;
