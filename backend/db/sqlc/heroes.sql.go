// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: heroes.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createHero = `-- name: CreateHero :one
INSERT INTO heroes(
  hero_name,
  format_id,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4
) RETURNING hero_id, hero_name, format_id, created_at, updated_at
`

type CreateHeroParams struct {
	HeroName  string        `json:"hero_name"`
	FormatID  uuid.NullUUID `json:"format_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (q *Queries) CreateHero(ctx context.Context, arg CreateHeroParams) (Hero, error) {
	row := q.queryRow(ctx, q.createHeroStmt, createHero,
		arg.HeroName,
		arg.FormatID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Hero
	err := row.Scan(
		&i.HeroID,
		&i.HeroName,
		&i.FormatID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteHero = `-- name: DeleteHero :exec
DELETE FROM heroes
WHERE hero_id = $1
`

func (q *Queries) DeleteHero(ctx context.Context, heroID uuid.UUID) error {
	_, err := q.exec(ctx, q.deleteHeroStmt, deleteHero, heroID)
	return err
}

const getAllHeroes = `-- name: GetAllHeroes :many
SELECT hero_id, hero_name, format_id, created_at, updated_at FROM HEROES
ORDER BY format_id
`

func (q *Queries) GetAllHeroes(ctx context.Context) ([]Hero, error) {
	rows, err := q.query(ctx, q.getAllHeroesStmt, getAllHeroes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Hero{}
	for rows.Next() {
		var i Hero
		if err := rows.Scan(
			&i.HeroID,
			&i.HeroName,
			&i.FormatID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHeroById = `-- name: GetHeroById :one
SELECT hero_id, hero_name, format_id, created_at, updated_at FROM heroes
WHERE hero_id = $1 LIMIT 1
`

func (q *Queries) GetHeroById(ctx context.Context, heroID uuid.UUID) (Hero, error) {
	row := q.queryRow(ctx, q.getHeroByIdStmt, getHeroById, heroID)
	var i Hero
	err := row.Scan(
		&i.HeroID,
		&i.HeroName,
		&i.FormatID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listHeroes = `-- name: ListHeroes :many
SELECT hero_id, hero_name, format_id, created_at, updated_at FROM heroes
ORDER BY hero_id
LIMIT $1
OFFSET $2
`

type ListHeroesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListHeroes(ctx context.Context, arg ListHeroesParams) ([]Hero, error) {
	rows, err := q.query(ctx, q.listHeroesStmt, listHeroes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Hero{}
	for rows.Next() {
		var i Hero
		if err := rows.Scan(
			&i.HeroID,
			&i.HeroName,
			&i.FormatID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHero = `-- name: UpdateHero :one
UPDATE heroes
SET
hero_name = coalesce($1, hero_name),
format_id = coalesce($2, format_id),
updated_at = coalesce($3, updated_at)
WHERE hero_id = $4
RETURNING hero_id, hero_name, format_id, created_at, updated_at
`

type UpdateHeroParams struct {
	HeroName  sql.NullString `json:"hero_name"`
	FormatID  uuid.NullUUID  `json:"format_id"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	HeroID    uuid.UUID      `json:"hero_id"`
}

func (q *Queries) UpdateHero(ctx context.Context, arg UpdateHeroParams) (Hero, error) {
	row := q.queryRow(ctx, q.updateHeroStmt, updateHero,
		arg.HeroName,
		arg.FormatID,
		arg.UpdatedAt,
		arg.HeroID,
	)
	var i Hero
	err := row.Scan(
		&i.HeroID,
		&i.HeroName,
		&i.FormatID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
