package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

type HeroModel struct {
	DB *db.Queries
}

func (m HeroModel) GetAllHeroes() ([]db.Hero, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	heroes, err := m.DB.GetAllHeroes(ctx)
	if err != nil {
		return nil, err
	}

	if heroes == nil {
		heroes = []db.Hero{}
	}

	return heroes, nil
}

func (m HeroModel) GetHeroById(id uuid.UUID) (*db.Hero, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hero, err := m.DB.GetHeroById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &hero, nil

}
