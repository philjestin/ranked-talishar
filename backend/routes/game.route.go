package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
)

type GameRoutes struct {
	gameController controllers.GameController
}

func NewRouteGame(gameController controllers.GameController) GameRoutes {
	return GameRoutes{gameController}
}

func (cr *GameRoutes) GameRoute(rg *gin.RouterGroup) {
	router := rg.Group("games")
	router.POST("/", cr.gameController.CreateGame)
	router.GET("/", cr.gameController.GetAllGames)
	router.PATCH("/:gameId", cr.gameController.UpdateGame)
	router.GET("/:gameId", cr.gameController.GetGameById)
	router.DELETE("/:gameId", cr.gameController.DeleteGameById)
}