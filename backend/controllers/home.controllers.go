package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	gintemplrenderer "github.com/philjestin/ranked-talishar/gintemplaterenderer"
	"github.com/philjestin/ranked-talishar/schemas"
	"github.com/philjestin/ranked-talishar/views"
)

type HomeController struct {
	db  *db.Queries
	ctx context.Context
}

func NewHomeController(db *db.Queries, ctx context.Context) *HomeController {
	return &HomeController{db, ctx}
}

func (cc *HomeController) Home() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		currentUser := ctx.Value("authorization_payload")

		fmt.Println(currentUser, " current-user / authorization_payload")

		user, err := cc.db.GetUserById(ctx, uuid.MustParse("90e8a5ec-da6a-4828-a095-b1c0555ad049"))
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Failed to retrieve user with this ID"})
				return
			}
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "Failed retrieving user", "error": err.Error()})
			return
		}

		rsp := schemas.CreateUserResponse{
			UserName:          user.UserName,
			UserEmail:         user.UserEmail,
			CreatedAt:         user.CreatedAt,
			PasswordChangedAt: user.PasswordChangedAt.Time,
		}
		response := gintemplrenderer.New(ctx.Request.Context(), http.StatusOK, views.Home(rsp))
		ctx.Render(http.StatusOK, response)
	}
}
