package elo

import (
	"context"
	"database/sql"
	"errors"
	"math"

	"github.com/google/uuid"

	db "github.com/philjestin/ranked-talishar/db/sqlc"
)

const (
  KFactor = 32  // K-factor for rating adjustments
  cValue  = 400 // Constant for expected score calculation
)

// Player represents a player with their current ELO rating
type Player struct {
  UserID uuid.NullUUID `json:"user_id"`
  Elo int32    `json:"elo"`
}


func UpdateRatings(ctx context.Context, q *db.Queries, playerAID, playerBID uuid.UUID, scoreA float64) error {
	// Fetch player information from database using IDs
	playerA, err := q.GetUserById(ctx, playerAID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}
	playerB, err := q.GetUserById(ctx, playerBID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}

	// Calculate expected scores based on current ratings
	expectedA := ExpectedScore(playerA.Elo, playerB.Elo)
	expectedB := 1 - expectedA

	// Perform calculations with float64 for better precision
	ratingChangeA := KFactor * (scoreA - expectedA)
	ratingChangeB := KFactor * ((1 - scoreA) - expectedB)

	// Round the rating changes to nearest integer before updating
	newRatingA := playerA.Elo + int32(math.Round(ratingChangeA))
	newRatingB := playerB.Elo + int32(math.Round(ratingChangeB))

  // Update player ratings in the database (using sqlc queries)
  playerAParams := db.UpdatePlayerRatingParams{
    Elo:    sql.NullInt32{Int32: int32(newRatingA), Valid: true},
    UserID: uuid.NullUUID{UUID: playerA.UserID, Valid: true},
  }
  err = q.UpdatePlayerRating(ctx, playerAParams)
  if err != nil {
    return err
  }

  playerBParams := db.UpdatePlayerRatingParams{
    Elo:    sql.NullInt32{Int32: int32(newRatingB), Valid: true},
    UserID: uuid.NullUUID{UUID: playerA.UserID, Valid: true},
  }
  err = q.UpdatePlayerRating(ctx, playerBParams)
  if err != nil {
    return err
  }

	return nil
}

// ExpectedScore calculates the expected score (win probability) for a player
func ExpectedScore(ratingA, ratingB int32) float64 {
	return 1 / (1 + math.Pow(10, float64(ratingB-ratingA)/cValue))
}
