-- name: CreateHero :one
INSERT INTO heroes(
  hero_name,
  format_id,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetHeroById :one
SELECT * FROM heroes
WHERE hero_id = $1 LIMIT 1;

-- name: ListHeroes :many
SELECT * FROM heroes
ORDER BY hero_id
LIMIT $1
OFFSET $2;

-- name: UpdateHero :one
UPDATE heroes
SET
hero_name = coalesce(sqlc.narg('hero_name'), hero_name),
format_id = coalesce(sqlc.narg('format_id'), format_id),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE hero_id = sqlc.arg('hero_id')
RETURNING *;

-- name: DeleteHero :exec
DELETE FROM heroes
WHERE hero_id = $1;

-- name: GetAllHeroes :many
SELECT * FROM HEROES
ORDER BY format_id;

-- name: GetHeroesByFormatId :many
SELECT * FROM heroes
where format_id = $1;