package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/schemas"
	"github.com/philjestin/ranked-talishar/token"
)

type RefreshController struct {
	db            *db.Queries
	ctx           context.Context
	jwtMaker      token.Maker
	tokenDuration time.Duration
	secretKey     string
}

func NewRefreshController(db *db.Queries, ctx context.Context, jwtMaker token.Maker, tokenDuration time.Duration, secretKey string) *RefreshController {
	return &RefreshController{db, ctx, jwtMaker, tokenDuration, secretKey}
}

func (cc *RefreshController) Refresh(ctx *gin.Context) {
	var req schemas.RefreshTokenRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Invalid request body",
			"error":  err.Error(),
		})
	}

	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cc.secretKey), nil
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Issue with refresh token",
			"error":  err.Error(),
		})
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		refreshToken, err := cc.db.GetRefreshTokenByUserID(ctx, uuid.MustParse(req.UserID))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "Issue looking up user in refresh tokens",
				"error":  err.Error(),
			})
			return
		}
		user, err := cc.db.GetUserById(ctx, refreshToken.UserID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "Issue looking up user",
				"error":  err.Error(),
			})
			return
		}

		tokens, err := cc.jwtMaker.CreateToken(user.UserName, cc.tokenDuration)
		if err != nil {
			// Handle error (e.g., failed to create new tokens)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": "Failed to generate new tokens",
				"error":   err.Error(),
			})
			return
		}

		if refreshToken.ID != 0 {
			_, err = cc.db.UpdateRefreshToken(ctx, db.UpdateRefreshTokenParams{
				ID:           refreshToken.ID,
				RefreshToken: sql.NullString{String: tokens.RefreshToken, Valid: tokens.RefreshToken != ""},
			})
		} else {
			_, err = cc.db.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
				UserID:       user.UserID,
				RefreshToken: tokens.RefreshToken,
				Expiry:       time.Now().Add(time.Hour * 24),
			})
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status":  "failed",
				"message": "Failed to update refresh token",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"accessToken":  tokens.AccessToken,
			"refreshToken": tokens.RefreshToken,
		})
	}
}
