package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	db "github.com/philjestin/ranked-talishar/db/sqlc"
	"github.com/philjestin/ranked-talishar/password"
	"github.com/philjestin/ranked-talishar/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
    db  *db.Queries
    ctx context.Context
}

func NewUserController(db *db.Queries, ctx context.Context) *UserController {
    return &UserController{db, ctx}
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
        ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Password failed", "error": err.Error()})
    }

    now := time.Now()
    args := &db.CreateUserParams{
        UserName:   payload.UserName,
        UserEmail:    payload.UserEmail,
        CreatedAt:   now,
        UpdatedAt:   now,
        HashedPassword: hashedPassword,
    }

    user, err := cc.db.CreateUser(ctx, *args)

    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "successfully created user", "user": user})
}

// Update user handler
func (cc *UserController) UpdateUser(ctx *gin.Context) {
    var payload *schemas.UpdateUser
    userId := ctx.Param("userId")

    if err := ctx.ShouldBindJSON(&payload); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
        return
    }

    now := time.Now()
    args := &db.UpdateUserParams{
        UserID:   uuid.MustParse(userId),
        UserName:   sql.NullString{String: payload.UserName, Valid: payload.UserName != ""},
        UserEmail:    sql.NullString{String: payload.UserEmail, Valid: payload.UserEmail != ""},
        UpdatedAt:   sql.NullTime{Time: now, Valid: true},
        HashedPassword: sql.NullString{String: payload.Password, Valid: payload.Password != ""},
    }

    user, err := cc.db.UpdateUser(ctx, *args)

    if err != nil {
        if err == sql.ErrNoRows {
            ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve user with this ID"})
            return
        }
        ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "successfully updated user", "user": user})
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

    ctx.JSON(http.StatusNoContent, gin.H{"status": "successfuly deleted"})

}