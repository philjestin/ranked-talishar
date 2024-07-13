package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
)

type HeroRoutes struct {
	heroController controllers.HeroController
}

func NewRouteHero(heroController controllers.HeroController) HeroRoutes {
	return HeroRoutes{heroController}
}

func (cr *HeroRoutes) HeroRoute(rg *gin.RouterGroup) {
	router := rg.Group("heroes")
	router.POST("/", cr.heroController.CreateHero)
	router.GET("/", cr.heroController.GetAllHeroes)
	router.PATCH("/:heroId", cr.heroController.UpdateHero)
	router.GET("/:heroId", cr.heroController.GetHeroById)
	router.DELETE("/:heroId", cr.heroController.DeleteHeroById)
}
