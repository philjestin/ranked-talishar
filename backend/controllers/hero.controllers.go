package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/schemas"

	"github.com/gin-gonic/gin"
)

type HeroController struct {
	db *db.Queries
	ctx context.Context
}

func NewHeroController(db *db.Queries, ctx context.Context) *HeroController {
	return &HeroController{db, ctx}
}

// Create hero handler
func (cc *HeroController) CreateHero(ctx *gin.Context) {
	var payload *schemas.CreateHero

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error": err.Error(),
		})
		return
	}

	now := time.Now()
	args := &db.CreateHeroParams{
		HeroName:   payload.HeroName,
		FormatID:    payload.FormatID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	hero, err := cc.db.CreateHero(ctx, *args)

	if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving hero", "error": err.Error()})
			return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "successfully created hero",
		"hero": hero,
	})
}

// Update Hero handler
func (cc *HeroController) UpdateHero(ctx *gin.Context) {
	var payload *schemas.UpdateHero
	heroId := ctx.Param("heroId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error": err.Error(),
		})
	}

	now := time.Now()
	args := &db.UpdateHeroParams{
		HeroID: 				  uuid.MustParse(heroId),
		HeroName: 				sql.NullString{String: payload.HeroName, Valid: payload.HeroName != ""},
		FormatID:  				payload.FormatID,
		UpdatedAt: 				sql.NullTime{Time: now, Valid: true},
	}

	hero, err := cc.db.UpdateHero(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "failed",
				"message": "Failed to retrieve hero with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving hero",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully updated hero",
		"hero": hero,
	})
}

// Get a single hero handler
func (cc *HeroController) GetHeroById(ctx *gin.Context) {
  heroId := ctx.Param("heroId")

	hero, err := cc.db.GetHeroById(ctx, uuid.MustParse(heroId))
	if err != nil {
		if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": "failed",
					"message": "Failed to retrieve hero with this ID",
				})
				return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving hero",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully retrieved id",
		"hero": hero,
	})
}

// Retrieve all heros handler
func (cc *HeroController) GetAllHeroes(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListHeroesParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	heros, err := cc.db.ListHeroes(ctx, *args)
	if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status": "Failed to retrieve heros",
				"error": err.Error(),
			})
			return
	}

	if heros == nil {
		heros = []db.Hero{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully retrieved all heros",
		"size": len(heros),
		"heros": heros,
	})
}

// Deleting Hero handler
func (cc *HeroController) DeleteHeroById(ctx *gin.Context) {
    heroId := ctx.Param("heroId")

    _, err := cc.db.GetHeroById(ctx, uuid.MustParse(heroId))
    if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": "failed",
					"message": "Failed to retrieve hero with this ID",
				})
				return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving hero",
			"error": err.Error(),
		})
		return
	}

	err = cc.db.DeleteHero(ctx, uuid.MustParse(heroId))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "failed",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "successfully deleted"})
}