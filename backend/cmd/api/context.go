package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")

func ContextSetUser(ctx *gin.Context, user *data.User) {
	ctx.Set(string(userContextKey), user)
}

func ContextGetUser(ctx *gin.Context) *data.User {
	user, ok := ctx.Value(userContextKey).(*data.User)

	if !ok {
		panic("missing user value in request context")
	}

	return user
}
