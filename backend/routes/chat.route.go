package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/philjestin/ranked-talishar/chat"
)

type ChatRoutes struct {
	hub *chat.Hub
}

func NewRouteChat(hub *chat.Hub) ChatRoutes {
	return ChatRoutes{hub}
}

func (cr *ChatRoutes) ChatRoute(rg *gin.RouterGroup) {
	router := rg.Group("ws")
	router.GET("/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		log.Printf("room from c: %v", roomId)
		chat.ServeWs(c, cr.hub)
	})
}
