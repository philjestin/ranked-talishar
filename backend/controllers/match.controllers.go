package controllers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/schemas"

	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/util"
)

type MatchController struct {
	db  *db.Queries
	ctx context.Context
}

func NewMatchController(db *db.Queries, ctx context.Context) *MatchController {
	return &MatchController{db, ctx}
}

// Create match handler
func (cc *MatchController) CreateMatch(ctx *gin.Context) {
	var payload *schemas.CreateMatch

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error":  err.Error(),
		})
		return
	}

	now := time.Now()
	args := &db.CreateMatchParams{
		MatchName: sql.NullString{String: payload.MatchName, Valid: payload.MatchName != ""},
		// Handle FormatID
		FormatID: func() uuid.NullUUID {
			if payload.FormatID != (uuid.UUID{}) { // Check if payload.FormatID is a zero-value UUID
				return uuid.NullUUID{UUID: payload.FormatID, Valid: true}
			}
			return uuid.NullUUID{Valid: false} // Set Valid to false if payload.FormatID is empty
		}(),

		Player1ID: func() uuid.NullUUID {
			if payload.Player1ID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player1ID, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),
		// Handle Player1Hero (similar to FormatID)
		Player1Hero: func() uuid.NullUUID {
			if payload.Player1Hero != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player1Hero, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),
		// Handle GameID
		GameID: func() uuid.NullUUID {
			if payload.GameID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.GameID, Valid: true}
			}
			return uuid.NullUUID{Valid: false} // Set Valid to false if empty
		}(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	match, err := cc.db.CreateMatch(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving match", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "successfully created match",
		"match":  match,
	})
}

// Update Match handler
func (cc *MatchController) UpdateMatch(ctx *gin.Context) {
	log.Println("Inside of the Match controller")
	var payload *schemas.UpdateMatch
	matchId := ctx.Param("matchId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error":  err.Error(),
		})
	}

	now := time.Now()
	args := &db.UpdateMatchParams{
		MatchID:   uuid.MustParse(matchId),
		MatchName: sql.NullString{String: payload.MatchName, Valid: payload.MatchName != ""},

		FormatID: func() uuid.NullUUID {
			if payload.FormatID != (uuid.UUID{}) { // Check if payload.FormatID is a zero-value UUID
				return uuid.NullUUID{UUID: payload.FormatID, Valid: true}
			}
			return uuid.NullUUID{Valid: false} // Set Valid to false if payload.FormatID is empty
		}(),

		UpdatedAt: sql.NullTime{Time: now, Valid: true},

		Player1ID: func() uuid.NullUUID {
			if payload.Player1ID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player1ID, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),

		// Handle Player1Hero (similar to FormatID)
		Player1Hero: func() uuid.NullUUID {
			if payload.Player1Hero != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player1Hero, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),

		Player2ID: func() uuid.NullUUID {
			if payload.Player2ID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player2ID, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),

		// Handle Player2Hero (similar to FormatID)
		Player2Hero: func() uuid.NullUUID {
			if payload.Player2Hero != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.Player2Hero, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),

		Player2Decklist: sql.NullString{String: payload.Player2Decklist, Valid: payload.Player2Decklist != ""},
		Player1Decklist: sql.NullString{String: payload.Player1Decklist, Valid: payload.Player1Decklist != ""},
		InProgress:      sql.NullBool{Valid: true, Bool: payload.InProgress},

		WinnerID: func() uuid.NullUUID {
			if payload.WinnerID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.WinnerID, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),

		LoserID: func() uuid.NullUUID {
			if payload.LoserID != (uuid.UUID{}) {
				return uuid.NullUUID{UUID: payload.LoserID, Valid: true}
			}
			return uuid.NullUUID{Valid: false}
		}(),
	}

	match, err := cc.db.UpdateMatch(ctx, *args)

	// Check if loser_id or winner_id has changed
	hasWinnerIDChanged := match.WinnerID.Valid
	hasLoserIDChanged := match.LoserID.Valid

	// Send notification only if winner and loser ID has changed
	if hasWinnerIDChanged && hasLoserIDChanged {
		err := util.SendMatchUpdateNotification(ctx, cc.db, match.MatchID, payload.WinnerID, payload.LoserID)
		if err != nil {
			log.Printf("Error sending notification for match update: %v\n", err)
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "Failed to retrieve match with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving match",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully updated match",
		"match":  match,
	})
}

// Get a single match handler
func (cc *MatchController) GetMatchById(ctx *gin.Context) {
	matchId := ctx.Param("matchId")

	match, err := cc.db.GetMatchById(ctx, uuid.MustParse(matchId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "Failed to retrieve match with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving match",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully retrieved id",
		"match":  match,
	})
}

// Retrieve all matches handler
func (cc *MatchController) GetAllMatches(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListMatchesParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	matches, err := cc.db.ListMatches(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed to retrieve matches",
			"error":  err.Error(),
		})
		return
	}

	if matches == nil {
		matches = []db.Match{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Successfully retrieved all matches",
		"size":    len(matches),
		"matches": matches,
	})
}

// Deleting Match handler
func (cc *MatchController) DeleteMatchById(ctx *gin.Context) {
	matchId := ctx.Param("matchId")

	_, err := cc.db.GetMatchById(ctx, uuid.MustParse(matchId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status":  "failed",
				"message": "Failed to retrieve match with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving match",
			"error":  err.Error(),
		})
		return
	}

	err = cc.db.DeleteMatch(ctx, uuid.MustParse(matchId))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "successfully deleted"})
}
