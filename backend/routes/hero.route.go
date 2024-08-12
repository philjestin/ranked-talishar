package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
	"github.com/philjestin/ranked-talishar/token"
)

type HeroRoutes struct {
	heroController controllers.HeroController
	jwtMaker       token.Maker
}

func NewRouteHero(heroController controllers.HeroController, jwtMaker token.Maker) HeroRoutes {
	return HeroRoutes{heroController, jwtMaker}
}

func (cr *HeroRoutes) HeroRoute(rg *gin.RouterGroup) {
	router := rg.Group("heroes")
	router.POST("/", cr.heroController.CreateHero)
	router.GET("/", cr.heroController.GetAllHeroes)
	router.PATCH("/:heroId", cr.heroController.UpdateHero)
	router.GET("/:heroId", cr.heroController.GetHeroById)
	router.DELETE("/:heroId", cr.heroController.DeleteHeroById)

}
