package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/schemas"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	db  *db.Queries
	ctx context.Context
}

type CreateGameArgs struct {
  GameName string `json:"game_name"`
}

func NewGameController(db *db.Queries, ctx context.Context) *GameController {
	return &GameController{db, ctx}
}

// Create game handler
func (cc *GameController) CreateGame(ctx *gin.Context) {
	var payload *schemas.CreateGame

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	args := &CreateGameArgs{
		GameName: payload.GameName,
	}

	game, err := cc.db.CreateGame(ctx, args.GameName)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving game", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully created game", "game": game})
}

// Update Game Handler
func (cc *GameController) UpdateGame(ctx *gin.Context) {
	var payload *schemas.UpdateGame
	gameId := ctx.Param("gameId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	args := db.UpdateGameParams{
		GameID: uuid.MustParse(gameId),
		GameName: sql.NullString{String: payload.GameName, Valid: payload.GameName != ""},
	}

	game, err := cc.db.UpdateGame(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve game with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving game", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated game", "game": game})	
}

// Get a single Game
func (cc *GameController) GetGameById(ctx *gin.Context) {
	gameId := ctx.Param("gameId")

	game, err := cc.db.GetGameByID(ctx, uuid.MustParse(gameId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "Failed", "message": "Failed to retrieve game with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving game", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved ID", "game": game})
}

// Retrieve all game handler
func (cc *GameController) GetAllGames(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID -1) * reqLimit

	args := &db.ListGamesParams{
		Limit: int32(reqLimit),
		Offset: int32(offset),
	}

	games, err := cc.db.ListGames(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed to retrieve games",
			"error": err.Error(),
		})
		return
	}

	if games == nil {
		games = []db.Game{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully retrieved all games",
		"size": len(games),
		"games": games,
	})
}

// Deleting game handler
func (cc *GameController) DeleteGameById(ctx *gin.Context) {
	gameId := ctx.Param("gameId")

	_, err := cc.db.GetGameByID(ctx, uuid.MustParse(gameId))

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "failed",
				"message": "Failed to retrieve game with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving game",
			"error": err.Error(),
		})
		return
	}
}