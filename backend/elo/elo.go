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
	ScalingFactor
)

// Player represents a player with their current ELO rating
type Player struct {
	UserID uuid.NullUUID `json:"user_id"`
	Elo    int32         `json:"elo"`
}

func UpdateRatings(ctx context.Context, q *db.Queries, winnerID, loserID uuid.UUID) error {
	// Fetch player information from database using IDs
	winner, err := q.GetUserById(ctx, winnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}
	loser, err := q.GetUserById(ctx, loserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("player not found")
		}
		return err
	}

	// Calculate rating difference
	ratingDiff := winner.Elo - loser.Elo

	// Adjust K-Factor based on rating difference
	KFactor := KFactor - int32(math.Abs(float64(ratingDiff))/ScalingFactor)

	// // Calculate expected scores based on current ratings
	expectedWinnerScore := ExpectedScore(winner.Elo, loser.Elo)
	expectedLoserScore := 1 - expectedWinnerScore

	// Perform calculations with float64 for better precision
	winnerRatingChange := float64(KFactor) * (1.0 - expectedWinnerScore)
	loserRatingChange := float64(KFactor) * (expectedLoserScore - 0)

	// Round the rating changes to nearest integer before updating
	newWinnerRating := winner.Elo + int32(math.Round(winnerRatingChange))
	newLoserRating := loser.Elo - int32(math.Round(loserRatingChange))

	// Update player ratings in the database (using sqlc queries)
	winnerParams := db.UpdatePlayerRatingParams{
		Elo:    sql.NullInt32{Int32: int32(newWinnerRating), Valid: true},
		UserID: uuid.NullUUID{UUID: winner.UserID, Valid: true},
	}
	err = q.UpdatePlayerRating(ctx, winnerParams)
	if err != nil {
		return err
	}

	loserParams := db.UpdatePlayerRatingParams{
		Elo:    sql.NullInt32{Int32: int32(newLoserRating), Valid: true},
		UserID: uuid.NullUUID{UUID: loser.UserID, Valid: true},
	}
	err = q.UpdatePlayerRating(ctx, loserParams)
	if err != nil {
		return err
	}

	return nil
}

// ExpectedScore calculates the expected score (win probability) for a player
func ExpectedScore(ratingA, ratingB int32) float64 {
	return 1 / (1 + math.Exp(float64(ratingB-ratingA)/400))
}
