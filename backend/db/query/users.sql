-- name: CreateUser :one
INSERT INTO users(
    user_name,
    user_email,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
user_name = coalesce(sqlc.narg('user_name'), user_name),
user_email = coalesce(sqlc.narg('user_email'), user_email),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE user_id = sqlc.arg('user_id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;