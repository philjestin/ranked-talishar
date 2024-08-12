package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/password"
	"github.com/philjestin/ranked-talishar/token"
)

type LoginForm struct {
	User     string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginController struct {
	db            *db.Queries
	ctx           context.Context
	jwtMaker      token.Maker
	tokenDuration time.Duration
}

func NewLoginController(db *db.Queries, ctx context.Context, jwtMaker token.Maker, tokenDuration time.Duration) *LoginController {
	return &LoginController{db, ctx, jwtMaker, tokenDuration}
}

func (cc *LoginController) UserLogin() gin.HandlerFunc {

	fmt.Println("before return func")

	return func(ctx *gin.Context) {

		username := ctx.PostForm("username")
		userPassword := ctx.PostForm("password")

		user, err := cc.db.GetUser(ctx, username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to find user with this username"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Failed retrieving user", "error": err.Error()})
			return
		}

		err = password.CheckPassword(userPassword, user.HashedPassword)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "Password does not match.",
				"error":   err.Error(),
			})
			return
		}

		tokens, err := cc.jwtMaker.CreateToken(user.UserName, cc.tokenDuration)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed",
				"error":  err.Error(),
			})
			return
		}

		expiry := time.Now().Add(time.Hour * 24)
		refreshTokenArgs := &db.CreateRefreshTokenParams{
			RefreshToken: tokens.RefreshToken,
			UserID:       user.UserID,
			Expiry:       expiry,
		}

		_, err = cc.db.CreateRefreshToken(ctx, *refreshTokenArgs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "Failed",
				"error":  err.Error(),
			})
			return
		}

		_, err = ctx.Cookie("ranked_talishar_cookie")

		if err != nil {
			ctx.SetCookie("ranked_talishar_cookie", tokens.AccessToken, 6400, "/", "localhost", false, true)
		}
		ctx.Header("authorization", tokens.AccessToken)
		ctx.Header("bearer", tokens.AccessToken)

		ctx.Request.Header.Set("authorization", tokens.AccessToken)
		ctx.Request.Header.Set("bearer", tokens.AccessToken)
		// authHeader = ctx.Head

		testHeader := ctx.GetHeader("authorization")

		fmt.Println("set cookie redirecting now")

		fmt.Println(testHeader)

		ctx.Redirect(http.StatusFound, "/home?fresh=true")
	}
}
