-- name: CreateRefreshToken :one
INSERT into refresh_tokens(
    user_id,
    refresh_token,
    expiry
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetRefreshTokenByUserID :one
SELECT * FROM refresh_tokens
WHERE user_id = $1 LIMIT 1;

-- name: UpdateRefreshToken :one
UPDATE refresh_tokens
SET
refresh_token = coalesce(sqlc.narg('refresh_token'), refresh_token),
expiry = coalesce(sqlc.narg('expiry'), expiry)
WHERE id = sqlc.arg('id')
returning *;