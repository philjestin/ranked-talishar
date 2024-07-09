package player

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

// Fetch player information from database using IDs
func UpdatePlayersWinLossColumns(ctx context.Context, q *db.Queries, winnerId, loserId uuid.UUID) error {
	winner, err := q.GetUserById(ctx, winnerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}
	
	loser, err := q.GetUserById(ctx, loserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}

	winnerParams := db.IncrementWinsParams{
		UserID: winner.UserID,
	}
	err = q.IncrementWins(ctx, winnerParams)
  if err != nil {
    return err
  }

	loserParams := db.IncrementLossesParams{
		UserID: loser.UserID,
	}
	err = q.IncrementLosses(ctx, loserParams)
  if err != nil {
    return err
  }
	return nil
}