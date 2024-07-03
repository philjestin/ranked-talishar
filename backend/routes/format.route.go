package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
)

type FormatRoutes struct {
	formatController controllers.FormatController
}

func NewRouteFormat(formatController controllers.FormatController) FormatRoutes {
	return FormatRoutes{formatController}
}

func (cr *FormatRoutes) FormatRoute(rg *gin.RouterGroup) {
	router := rg.Group("formats")
	router.POST("/", cr.formatController.CreateFormat)
	router.GET("/", cr.formatController.GetAllFormats)
	router.PATCH("/:formatId", cr.formatController.UpdateFormat)
	router.GET("/:formatId", cr.formatController.GetFormatById)
	router.DELETE("/:formatId", cr.formatController.DeleteFormatById)
}