package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Request: ", ctx.Request)
		fmt.Println("Request Headers: ", ctx.Request.Header)
		fmt.Println("Request Body: ", ctx.Request.Body)

		ctx.Next()
	}
}
