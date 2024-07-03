-- name: CreateGame :one
INSERT INTO games(
  game_name
) VALUES (
  $1
) RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games
WHERE game_id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM games
ORDER BY game_id
LIMIT $1
OFFSET $2;

-- name: UpdateGame :one
UPDATE games
SET
game_name = COALESCE(sqlc.narg('game_name'), game_name)
WHERE game_id = sqlc.arg('game_id')
RETURNING *;

-- name: DeleteGame :exec
DELETE FROM games
WHERE game_id = $1;
