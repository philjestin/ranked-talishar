package schemas

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type CreateMatch struct {
    GameID          uuid.UUID `json:"game_id" binding:"required"`
    FormatID        uuid.UUID `json:"format_id" binding:"required"`
    MatchDate       time.Time `json:"match_date"`
    MatchName       string    `json:"match_name" binding:"required"`
    Player1ID       uuid.UUID `json:"player1_id" binding:"required"`
    Player2ID       uuid.UUID `json:"player2_id"`
    Player1Decklist string    `json:"player1_decklist" binding:"required"`
    Player2Decklist string    `json:"player2_decklist"`
    Player1Hero     uuid.UUID `json:"player1_hero" binding:"required"`
    Player2Hero     uuid.UUID `json:"player2_hero"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

type UpdateMatch struct {
    GameID          uuid.UUID    `json:"game_id"`
    FormatID        uuid.UUID    `json:"format_id"`
    MatchDate       sql.NullTime `json:"match_date"`
    MatchName       string       `json:"match_name"`
    Player1Decklist string       `json:"player1_decklist"`
    Player2Decklist string       `json:"player2_decklist"`
    Player1Hero     uuid.UUID    `json:"player1_hero"`
    Player2Hero     uuid.UUID    `json:"player2_hero"`
    UpdatedAt       sql.NullTime `json:"updated_at"`
    InProgress      bool         `json:"in_progress"`
		WinnerID				uuid.UUID		 `json:"winner_id"`
		LoserID					uuid.UUID		 `json:"loser_id"`
		Player1ID       uuid.UUID `json:"player1_id"`
    Player2ID       uuid.UUID `json:"player2_id"`
}