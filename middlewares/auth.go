package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/EventManagement-Golang/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Request"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	ctx.Set("userId", userId)
	ctx.Next()
}
