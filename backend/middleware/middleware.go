package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/philjestin/ranked-talishar/token"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	RefreshHeaderKey        = "X-Refresh-Token"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("Inside auth middleware")

		authorizationHeader, _ := ctx.Cookie("ranked_talishar_cookie")
		// fmt.Println(cookie)

		// authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		fmt.Println(authorizationHeader, "authorization header")

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			handleUnauthorized(ctx, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		fmt.Println("fields", fields)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			handleUnauthorized(ctx, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			handleUnauthorized(ctx, err)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil {
			handleUnauthorized(ctx, err)
			return
		}

		fmt.Println("accesstoken", accessToken)

		fmt.Println("payload ", payload)

		// var userModel *data.UserModel
		// user, err := userModel.GetForToken(accessToken)
		// if err != nil {
		// 	handleUnauthorized(ctx, err)
		// }

		// ctx.Set("user", user)

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}

func handleUnauthorized(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status":  "failed",
		"message": "Unauthorized",
		"error":   err.Error(),
	})
}
