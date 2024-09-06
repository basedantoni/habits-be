-- name: GetUser :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email
LIMIT 20;

-- name: CreateUser :one
INSERT INTO users (
  id, email, password, created_at, updated_at
) VALUES (
  ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  SET updated_at = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
