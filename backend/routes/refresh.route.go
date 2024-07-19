package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
)

type RefreshRoutes struct {
	refreshController controllers.RefreshController
}

func NewRouteRefresh(refreshController controllers.RefreshController) RefreshRoutes {
	return RefreshRoutes{refreshController}
}

func (cr *RefreshRoutes) RefreshRoute(rg *gin.RouterGroup) {
	router := rg.Group("refresh")
	router.POST("/", cr.refreshController.Refresh)
}
