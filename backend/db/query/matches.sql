-- name: CreateMatch :one
INSERT INTO matches(
  match_id,
  game_id,
  format_id,
  match_date,
  match_name,
  player1_id,
  player2_id,
  player1_decklist,
  player2_decklist,
  player1_hero,
  player2_hero,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetMatchById :one
SELECT * FROM matches
WHERE match_id = $1 LIMIT 1;

-- name: ListMatches :many
SELECT * FROM matches
ORDER BY match_id
LIMIT $1
OFFSET $2;

-- name: UpdateMatch :one
UPDATE matches
SET
game_id = coalesce(sqlc.narg('game_id'), game_id),
format_id = coalesce(sqlc.narg('format_id'), format_id),
match_date = coalesce(sqlc.narg('match_date'), match_date),
match_name = COALESCE(sqlc.narg('match_name'), match_name),
player1_decklist = COALESCE(sqlc.narg('player1_decklist'), player1_decklist),
player2_decklist = COALESCE(sqlc.narg('player2_decklist'), player2_decklist),
player1_hero = COALESCE(sqlc.narg('player1_hero'), player1_hero),
player2_hero = COALESCE(sqlc.narg('player2_hero'), player2_hero),
updated_at = COALESCE(sqlc.narg('updated_at'), updated_at),
in_progress = COALESCE(sqlc.narg('in_progress'), in_progress),
winner_id = COALESCE(sqlc.narg('winner_id'), winner_id),
loser_id = COALESCE(sqlc.narg('loser_id'), loser_id),
player1_id = COALESCE(sqlc.narg('player1_id'), player1_id),
player2_id = COALESCE(sqlc.narg('player2_id'), player2_id)
WHERE match_id = sqlc.arg('match_id')
RETURNING *;

-- name: DeleteMatch :exec
DELETE FROM matches
WHERE match_id = $1;

-- name: GetMatchPlayers :many
SELECT winner_id, loser_id FROM matches where match_id = sqlc.arg('match_id');