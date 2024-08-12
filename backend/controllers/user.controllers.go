package controllers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/middleware"
	"github.com/philjestin/ranked-talishar/password"
	"github.com/philjestin/ranked-talishar/schemas"
	"github.com/philjestin/ranked-talishar/token"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	db            *db.Queries
	ctx           context.Context
	jwtMaker      token.Maker
	tokenDuration time.Duration
}

func NewUserController(db *db.Queries, ctx context.Context, jwtMaker token.Maker, tokenDuration time.Duration) *UserController {
	return &UserController{db, ctx, jwtMaker, tokenDuration}
}

func newUserResponse(user db.User) schemas.CreateUserResponse {
	return schemas.CreateUserResponse{
		UserName:          user.UserName,
		UserEmail:         user.UserEmail,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt.Time,
	}
}

func (cc *UserController) SignupUser(ctx *gin.Context) {
	var payload *schemas.SignupUserRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error":  err.Error(),
		})
		return
	}

	hashedPassword, err := password.HashedPassword((payload.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Password did not match criteria",
			"error":  err.Error(),
		})
	}

	now := time.Now()
	args := &db.CreateUserParams{
		UserName:       payload.UserName,
		UserEmail:      payload.UserEmail,
		CreatedAt:      now,
		UpdatedAt:      now,
		HashedPassword: hashedPassword,
	}

	user, err := cc.db.CreateUser(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed creating user",
			"error":  err.Error(),
		})
		return
	}

	tokens, err := cc.jwtMaker.CreateToken(user.UserName, cc.tokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Failed signing user in",
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

	rsp := schemas.LoginUserResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (cc *UserController) LoginUser(ctx *gin.Context) {
	var req schemas.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Failed payload",
			"error":  err.Error(),
		})
		return
	}

	user, err := cc.db.GetUser(ctx, req.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to find user with this username"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	err = password.CheckPassword(req.Password, user.HashedPassword)
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

	rsp := schemas.LoginUserResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		User:         newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

// Create user handler
func (cc *UserController) CreateUser(ctx *gin.Context) {
	var payload *schemas.CreateUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	hashedPassword, err := password.HashedPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Password failed", "error": err.Error()})
	}

	now := time.Now()
	args := &db.CreateUserParams{
		UserName:       payload.UserName,
		UserEmail:      payload.UserEmail,
		CreatedAt:      now,
		UpdatedAt:      now,
		HashedPassword: hashedPassword,
	}

	user, err := cc.db.CreateUser(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	rsp := newUserResponse(user)

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully created user", "user": rsp})
}

// Update user handler
func (cc *UserController) UpdateUser(ctx *gin.Context) {
	var payload *schemas.UpdateUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	user, err := cc.db.GetUser(ctx, payload.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to find user with this username"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	if user.UserName != authPayload.UserName {
		err := errors.New("account does not belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"message": "Unauthorized",
			"error":   err.Error(),
		})
		return
	}

	hashedPassword, err := password.HashedPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Password failed", "error": err.Error()})
		return
	}

	now := time.Now()
	args := &db.UpdateUserParams{
		UserID:         uuid.MustParse(payload.UserId),
		UserName:       sql.NullString{String: payload.UserName, Valid: payload.UserName != ""},
		UserEmail:      sql.NullString{String: payload.UserEmail, Valid: payload.UserEmail != ""},
		UpdatedAt:      sql.NullTime{Time: now, Valid: true},
		HashedPassword: sql.NullString{String: hashedPassword, Valid: hashedPassword != ""},
	}

	user, err = cc.db.UpdateUser(ctx, *args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve user with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	rsp := schemas.UpdateUserResponse{
		UserName:          user.UserName,
		UserEmail:         user.UserEmail,
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt.Time,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated user", "user": rsp})
}

// Get a single handler
func (cc *UserController) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	user, err := cc.db.GetUserById(ctx, uuid.MustParse(userId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve user with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved id", "user": user})
}

// Retrieve all records handlers
func (cc *UserController) GetAllUsers(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	reqPageID, _ := strconv.Atoi(page)
	reqLimit, _ := strconv.Atoi(limit)
	offset := (reqPageID - 1) * reqLimit

	args := &db.ListUsersParams{
		Limit:  int32(reqLimit),
		Offset: int32(offset),
	}

	users, err := cc.db.ListUsers(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed to retrieve users", "error": err.Error()})
		return
	}

	if users == nil {
		users = []db.User{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Successfully retrieved all users", "size": len(users), "users": users})
}

// Deleting user handlers
func (cc *UserController) DeleteUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	_, err := cc.db.GetUserById(ctx, uuid.MustParse(userId))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve user with this ID"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
		return
	}

	err = cc.db.DeleteUser(ctx, uuid.MustParse(userId))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "failed", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{"status": "successfully deleted"})

}
