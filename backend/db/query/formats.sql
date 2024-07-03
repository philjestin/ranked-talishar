-- name: CreateFormat :one
INSERT INTO formats(
    format_name,
    format_description,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetFormatById :one
SELECT * FROM formats
WHERE format_id = $1 LIMIT 1;

-- name: ListFormats :many
SELECT * FROM formats
ORDER BY format_id
LIMIT $1
OFFSET $2;

-- name: UpdateFormat :one
UPDATE formats
SET
format_name = coalesce(sqlc.narg('format_name'), format_name),
format_description = coalesce(sqlc.narg('format_description'), format_description),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE format_id = sqlc.arg('format_id')
RETURNING *;

-- name: DeleteFormat :exec
DELETE FROM formats
WHERE format_id = $1;