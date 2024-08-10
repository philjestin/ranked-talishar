package controllers

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	db "github.com/philjestin/ranked-talishar/db/sqlc"
	gintemplrenderer "github.com/philjestin/ranked-talishar/gintemplaterenderer"
	"github.com/philjestin/ranked-talishar/schemas"
	"github.com/philjestin/ranked-talishar/views"
)

type TempleHeroController struct {
	db  *db.Queries
	ctx context.Context
}

func NewTempleHeroController(db *db.Queries, ctx context.Context) *TempleHeroController {
	return &TempleHeroController{db, ctx}
}

func (cc *TempleHeroController) ViewHeros() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		data := make(chan schemas.SlotContents)

		// We know there are 1 slots, so start a WaitGround.
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			heroes, err := cc.db.GetAllHeroes(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"status": "Failed to retrieve heroes",
					"error":  err.Error(),
				})
				return
			}

			if heroes == nil {
				heroes = []db.Hero{}
				data <- schemas.SlotContents{
					Name:     "heroData",
					Contents: views.HeroData(heroes),
				}
			}

			data <- schemas.SlotContents{
				Name:     "heroData",
				Contents: views.HeroData(heroes),
			}
		}()

		go func() {
			wg.Wait()
			close(data)
		}()

		response := gintemplrenderer.New(ctx.Request.Context(), http.StatusOK, views.Heroes(data))
		ctx.Render(http.StatusOK, response)
	}
}
