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

type FormatController struct {
	db *db.Queries
	ctx context.Context
}

func NewFormatController(db *db.Queries, ctx context.Context) *FormatController {
	return &FormatController{db, ctx}
}

// Create format handler
func (cc *FormatController) CreateFormat(ctx *gin.Context) {
	var payload *schemas.CreateFormat

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error": err.Error(),
		})
		return
	}

	now := time.Now()
	args := &db.CreateFormatParams{
		FormatName:   payload.FormatName,
		FormatDescription:    payload.FormatDescription,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	contact, err := cc.db.CreateFormat(ctx, *args)

	if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving format", "error": err.Error()})
			return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "successfully created format",
		"contact": contact,
	})
}

// Update Format handler
func (cc *FormatController) UpdateFormat(ctx *gin.Context) {
	var payload *schemas.UpdateFormat
	formatId := ctx.Param("formatId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error": err.Error(),
		})
	}

	now := time.Now()
	args := &db.UpdateFormatParams{
		FormatID: 				  uuid.MustParse(formatId),
		FormatName: 				sql.NullString{String: payload.FormatName, Valid: payload.FormatName != ""},
		FormatDescription:  sql.NullString{String: payload.FormatDescription, Valid: payload.FormatDescription != ""},
		UpdatedAt: 					sql.NullTime{Time: now, Valid: true},
	}

	format, err := cc.db.UpdateFormat(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "failed",
				"message": "Failed to retrieve format with this ID",
			})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving format",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully updated format",
		"format": format,
	})
}

// Get a single format handler
func (cc *FormatController) GetFormatById(ctx *gin.Context) {
    formatId := ctx.Param("formatId")

    format, err := cc.db.GetFormatById(ctx, uuid.MustParse(formatId))
    if err != nil {
			if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, gin.H{
						"status": "failed",
						"message": "Failed to retrieve format with this ID",
					})
					return
			}
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status": "Failed retrieving format",
				"error": err.Error(),
			})
			return
    }

    ctx.JSON(http.StatusOK, gin.H{
			"status": "Successfully retrieved id",
			"format": format,
	})
}

// Retrieve all formats handler
func (cc *FormatController) GetAllFormats(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListFormatsParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	formats, err := cc.db.ListFormats(ctx, *args)
	if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"status": "Failed to retrieve formats",
				"error": err.Error(),
			})
			return
	}

	if formats == nil {
		formats = []db.Format{}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully retrieved all formats",
		"size": len(formats),
		"formats": formats,
	})
}

// Deleting Format handler
func (cc *FormatController) DeleteFormatById(ctx *gin.Context) {
    formatId := ctx.Param("formatId")

    _, err := cc.db.GetFormatById(ctx, uuid.MustParse(formatId))
    if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"status": "failed",
					"message": "Failed to retrieve format with this ID",
				})
				return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "Failed retrieving format",
			"error": err.Error(),
		})
		return
	}

	err = cc.db.DeleteFormat(ctx, uuid.MustParse(formatId))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status": "failed",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "successfully deleted"})
}