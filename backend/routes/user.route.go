package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/controllers"
	"github.com/philjestin/ranked-talishar/middleware"
	"github.com/philjestin/ranked-talishar/token"
)

type UserRoutes struct {
	userController controllers.UserController
	jwtMaker       token.Maker
}

func NewRouteUser(userController controllers.UserController, jwtMaker token.Maker) UserRoutes {
	return UserRoutes{userController, jwtMaker}
}

func (cr *UserRoutes) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")
	router.POST("/", cr.userController.CreateUser)
	router.POST("/login", cr.userController.LoginUser)
	router.POST("/signup", cr.userController.SignupUser)

	authRoutes := rg.Group("/").Use(middleware.AuthMiddleware(cr.jwtMaker))

	authRoutes.GET("/", cr.userController.GetAllUsers)
	authRoutes.PATCH("/:userId", cr.userController.UpdateUser)
	authRoutes.GET("/:userId", cr.userController.GetUserById)
	authRoutes.DELETE("/:userId", cr.userController.DeleteUserById)
}
