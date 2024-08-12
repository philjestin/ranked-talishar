-- name: CreateUser :one
INSERT INTO users(
    user_name,
    user_email,
    created_at,
    updated_at,
    hashed_password
) VALUES (
    $1, $2, $3, $4, $5
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
updated_at = coalesce(sqlc.narg('updated_at'), updated_at),
hashed_password = COALESCE(sqlc.narg('hashed_password'), hashed_password),
password_changed_at = COALESCE(sqlc.narg('password_changed_at'), password_changed_at)
WHERE user_id = sqlc.arg('user_id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: UpdatePlayerRating :exec
UPDATE users
SET
elo = COALESCE(sqlc.narg('elo'), elo)
where user_id = sqlc.narg('user_id');

-- name: IncrementWins :exec
UPDATE users
SET
wins = coalesce(sqlc.narg('wins'), wins + 1)
WHERE user_id = sqlc.arg('user_id');

-- name: IncrementLosses :exec
UPDATE users
SET
losses = coalesce(sqlc.narg('losses'), losses + 1)
WHERE user_id = sqlc.arg('user_id');

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1
LIMIT 1;

-- name: GetForToken :one
select users.user_id, users.created_at, users.user_email, users.user_name, users.hashed_password, users.password_changed_at
FROM users
INNER JOIN refresh_tokens
on users.user_id = refresh_tokens.user_id
WHERE refresh_tokens.refresh_token = $1;
