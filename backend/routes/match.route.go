package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
)

type MatchRoutes struct {
	matchController controllers.MatchController
}

func NewRouteMatch(matchController controllers.MatchController) MatchRoutes {
	return MatchRoutes{matchController}
}

func (cr *MatchRoutes) MatchRoute(rg *gin.RouterGroup) {
	router := rg.Group("matches")
	router.POST("/", cr.matchController.CreateMatch)
	router.GET("/", cr.matchController.GetAllMatches)
	router.PATCH("/:matchId", cr.matchController.UpdateMatch)
	router.GET("/:matchId", cr.matchController.GetMatchById)
	router.DELETE("/:matchId", cr.matchController.DeleteMatchById)
}